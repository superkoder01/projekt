package test

import (
	"NotificationEmailService/config"
	"NotificationEmailService/config/smtp_security"
	"NotificationEmailService/internal/logger"
	"NotificationEmailService/internal/server"
	"NotificationEmailService/pkg/email"
	"github.com/ory/dockertest/v3"
	"github.com/sandvikcode/mockserver-client-go/pkg/mockclient"
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	rabbitMqConfig "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	rabbitMqModel "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

var (
	client     email.Client
	mailHogUrl string

	serverConfig = config.EmailServiceConfig{
		Service: config.ServiceConfig{
			WorkerPoolSize:        1,
			ErrorEmailTo:          []string{"support0@ovoo.pl", "support1@ovoo.pl"},
			ErrorEmailEnvironment: "test",
		},
		Smtp: config.SmtpConfig{
			Host:       "localhost",
			User:       "info@chain4.energy",
			Password:   "St48ENgfZDgLjef",
			Security:   smtp_security.None,
			SenderName: "Chain For Energy Notification",
		},
		Rabbitmq: rabbitMqConfig.RabbitMQConsumerConfig{
			QueueName:     "email-que",
			QueueType:     "quorum",
			IsAutoAck:     false,
			PrefetchCount: 1,
			RabbitMQConfig: rabbitMqConfig.RabbitMQConfig{
				IsDurable:    true,
				IsAutoDelete: false,
			},
		},
		Conversion: conversion.Config{Timeout: 10},
	}

	clientCfg = email.Config{
		BindingKey: "email",
		RabbitMQ: rabbitMqConfig.RabbitMQProducerConfig{
			ExchangeName: "test",
			ExchangeType: "direct",
			QueueBindings: []rabbitMqConfig.QueueBinding{{
				BindingKey: "email",
				QueueName:  []string{"email-que"},
			},
			},
			RabbitMQConfig: rabbitMqConfig.RabbitMQConfig{
				IsDurable:    true,
				IsAutoDelete: false,
			},
		},
	}
)

func TestMain(m *testing.M) {
	log.Println("Initializing tests...")

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	var resources []*dockertest.Resource
	resources = append(resources, runRabbitMq(pool, &serverConfig, &clientCfg))
	resources = append(resources, runMockServer(pool, &serverConfig.Conversion))

	var mailHogResource *dockertest.Resource
	mailHogResource, mailHogUrl = runMailHog(pool, &serverConfig)
	resources = append(resources, mailHogResource)

	if !waitForRabbitMq(&serverConfig.Rabbitmq) {
		log.Fatalf("rabbitmq is not running")
	}
	if !waitForMockServer(&serverConfig.Conversion) {
		log.Fatalf("mockserver is not running")
	}

	go func() {
		if err := server.NewByConfig(&serverConfig).RunEmailServer(); err != nil {
			log.Fatalf("Cannot run server. err: %v", err)
		}
	}()

	client, err = email.New(logger.NewApiLogger(&config.Logger{}), &clientCfg)
	if err != nil {
		log.Fatalf("Could not create client. err: %v", err)
	}

	code := m.Run()

	for _, resource := range resources {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	os.Exit(code)
}

func TestSend(t *testing.T) {
	startTime := newTime(t)
	mockServer := newMockClient(t, &serverConfig)

	mockServer.AddExpectation(
		mockclient.CreateExpectation(
			mockclient.WhenRequestMethod("POST"),
			mockclient.WhenRequestPath("/api/convert-pdf"),
			mockclient.ThenResponseStatus(http.StatusCreated),
		))
	defer mockServer.Clear("/api/convert-pdf")

	send(t, client, &rabbitMqModel.Email{
		Destination: []string{"tomasz.pawlik@ovoo.pl"},
		Title:       "title",
		Body:        "body",
		Attachments: CreateAttachments("attachmentFileName", "data"),
	})

	messages := getMessages(t, mailHogUrl, startTime)
	assert.NotNil(t, messages)
	assert.Equal(t, "info@chain4.energy", messages.Raw.From)
	assert.Equal(t, "tomasz.pawlik@ovoo.pl", messages.Raw.To[0])
	assert.Equal(t, "title", messages.Content.Headers.Subject[0])
	assert.Equal(t, "body", messages.MIME.Parts[0].Body)
	assert.True(t, strings.Contains(messages.Content.Body, "\"attachmentFileName.pdf\""))
}

func TestConversionInternalError(t *testing.T) {
	startTime := newTime(t)
	mockServer := newMockClient(t, &serverConfig)

	conversionError, err := conversion.NewInternalErrorByDetails("test-error").ToString()
	assert.Nil(t, err)
	mockServer.AddExpectation(
		mockclient.CreateExpectation(
			mockclient.WhenRequestMethod("POST"),
			mockclient.WhenRequestPath("/api/convert-pdf"),
			mockclient.ThenResponseStatus(http.StatusInternalServerError),
			mockclient.ThenResponseJSON(conversionError),
		))
	defer mockServer.Clear("/api/convert-pdf")

	send(t, client, &rabbitMqModel.Email{
		Destination: []string{"tomasz.pawlik@ovoo.pl"},
		Title:       "title",
		Body:        "body",
		Attachments: CreateAttachments("attachmentFileName", "data"),
	})

	messages := getMessages(t, mailHogUrl, startTime)
	assert.NotNil(t, messages)
	assert.Equal(t, "info@chain4.energy", messages.Raw.From)
	assert.Equal(t, "support0@ovoo.pl", messages.Raw.To[0])
	assert.Equal(t, "support1@ovoo.pl", messages.Raw.To[1])
	assert.Equal(t, "Email notification service error", messages.Content.Headers.Subject[0])
	assert.Equal(t, "To: [tomasz.pawlik@ovoo.pl]\r\nCC: []\r\nBCC: []\r\nTitle: title\r\n:Body: body\r\nAttachments: map[attachmentFileName:\"data\"]\r\nError: Cannot generate file\r\nError details: test-error\r\nEnvironment: test", messages.Content.Body)
}

func TestConversionConnectionError(t *testing.T) {
	startTime := newTime(t)
	mockServer := newMockClient(t, &serverConfig)

	conversionError, err := conversion.NewConnectionErrorByDetails("test-error").ToString()
	assert.Nil(t, err)
	mockServer.AddExpectation(
		mockclient.CreateExpectation(
			mockclient.WhenRequestMethod("POST"),
			mockclient.WhenRequestPath("/api/convert-pdf"),
			mockclient.WhenTimes(1),
			mockclient.ThenResponseStatus(http.StatusInternalServerError),
			mockclient.ThenResponseJSON(conversionError),
		))
	mockServer.AddExpectation(
		mockclient.CreateExpectation(
			mockclient.WhenRequestMethod("POST"),
			mockclient.WhenRequestPath("/api/convert-pdf"),
			mockclient.WhenTimes(1),
			mockclient.ThenResponseStatus(http.StatusCreated),
		))
	defer mockServer.Clear("/api/convert-pdf")

	send(t, client, &rabbitMqModel.Email{
		Destination: []string{"tomasz.pawlik@ovoo.pl"},
		Title:       "title",
		Body:        "body",
		Attachments: CreateAttachments("attachmentFileName", "data"),
	})

	messages := getMessages(t, mailHogUrl, startTime)
	assert.NotNil(t, messages)
	assert.Equal(t, "info@chain4.energy", messages.Raw.From)
	assert.Equal(t, "tomasz.pawlik@ovoo.pl", messages.Raw.To[0])
	assert.Equal(t, "title", messages.Content.Headers.Subject[0])
	assert.Equal(t, "body", messages.MIME.Parts[0].Body)
	assert.True(t, strings.Contains(messages.Content.Body, "\"attachmentFileName.pdf\""))
}
