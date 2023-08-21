/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package core

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/logger/zap"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/messagebroker/rabbitmq/domain"
	"log"
)

type rabbitMQProducer struct {
	rabbitMQCore
	config config.RabbitMQProducerConfig
}

/*
NewRabbitMqProducer creates client to publish messages. Every client create exchange, opens amqp.Connection and amqp.Channel.
If connection with RabbitMQ cannot be established due to configuration / external errors function calls panic.
Client use internal console logger, to customize logs see NewRabbitMqProducerWithCustomLogger
*/
func NewRabbitMqProducer(mqConfig config.RabbitMQProducerConfig) *rabbitMQProducer {
	appLogger := zap.NewZapLogger(
		&config.LoggerConfig{
			Development: false,
			Encoding:    "console",
			Level:       "debug",
		})

	producer := rabbitMQProducer{
		config: mqConfig,
		rabbitMQCore: rabbitMQCore{
			logger: zap.CreateLoggerContext(*appLogger, "RabbitMQ-Producer"),
		},
	}
	producer.initializeConnection()
	return &producer
}

/*
NewRabbitMqProducerWithLogger creates client to publish messages. Every client create exchange, opens amqp.Connection and amqp.Channel.
If connection with RabbitMQ cannot be established due to configuration / external errors function calls panic.
Client use injected logger.Logger implementation
*/
func NewRabbitMqProducerWithLogger(mqConfig config.RabbitMQProducerConfig, logger logger.Logger) *rabbitMQProducer {
	producer := rabbitMQProducer{
		config: mqConfig,
		rabbitMQCore: rabbitMQCore{
			logger: logger,
		},
	}
	producer.initializeConnection()

	return &producer
}

func (rabbit *rabbitMQProducer) initializeConnection() {
	if len(rabbit.config.ExchangeName) < 1 {
		rabbit.logger.Fatalf("Invalid configuration !! Exchange must be declared")
		return
	}

	connErr := rabbit.rabbitMQCore.createConnectionWithRabbitMQ(rabbit.config.RabbitUrl)
	if connErr != nil {
		rabbit.logger.Fatalf("Failed to create connection with rabbitMq ! %v", connErr)
		return
	} else {
		rabbit.logger.Infof("Successfully connected to rabbitMQCore service")
	}

	exErr := rabbit.createExchange(
		rabbit.config.ExchangeName,
		rabbit.config.ExchangeType,
		rabbit.config.IsDurable,
		rabbit.config.IsAutoDelete,
	)
	if exErr != nil {
		rabbit.logger.Fatalf("Failed to create exchange ! %v", exErr)
		return
	}
	err := rabbit.createBindings(rabbit.config.QueueBindings, rabbit.config.ExchangeName)
	if err != nil {
		rabbit.logger.Fatalf("Failed to create exchange ! %v", exErr)
		return
	}
	rabbit.logger.Debugf("RabbitProducer initialized")
}

/*	Publish message to exchange with bindingKey */
func (rabbit *rabbitMQProducer) PublishMessage(message domain.Message, bindingKey string) error {
	amqpMessage := prepareMessage(message)

	err := rabbit.channel.Publish(
		rabbit.config.ExchangeName,
		bindingKey,
		false,
		false,
		*amqpMessage,
	)

	if err != nil {
		log.Printf("Canot publish message: %v error: %v", message, err)
		return err
	}

	if len(message.MessageId) > 1 {
		log.Printf("Message with id %v successfully send !", message.MessageId)
	} else {
		log.Printf("Message successfully send !")
	}

	return nil
}

func prepareMessage(message domain.Message) *amqp.Publishing {
	var messageId, contentType string
	var deliveryMode uint8

	if len(message.MessageId) > 1 {
		messageId = message.MessageId
	} else {
		messageId = uuid.New().String()
	}

	if len(message.ContentType) > 1 {
		contentType = message.ContentType
	} else {
		contentType = "application/json"
	}

	if message.IsTransient == true {
		deliveryMode = amqp.Transient
	} else {
		deliveryMode = amqp.Persistent
	}

	return &amqp.Publishing{
		Headers:      message.Headers,
		DeliveryMode: deliveryMode,
		ContentType:  contentType,
		Body:         message.Payload,
		MessageId:    messageId,
	}
}
