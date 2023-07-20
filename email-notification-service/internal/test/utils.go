package test

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/test/mailhog"
	"NotificationEmailService/pkg/email"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/sandvikcode/mockserver-client-go/pkg/mockclient"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	rabbitMqConfig "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	rabbitMqModel "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"log"
	"net/http"
	"testing"
	"time"
)

func waitForRabbitMq(cfg *rabbitMqConfig.RabbitMQConsumerConfig) bool {
	for i := 0; i < 200; i++ {
		connection, err := amqp.Dial(cfg.RabbitUrl)
		if err == nil {
			channel, err := connection.Channel()
			if err == nil {
				_, err = channel.QueueDeclare(
					cfg.QueueName,
					cfg.IsDurable,
					cfg.IsAutoDelete,
					false,
					false,
					map[string]interface{}{"x-queue-type": cfg.QueueType},
				)
				if err == nil {
					return true
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	return false
}

func runRabbitMq(pool *dockertest.Pool, serverConfig *config.EmailServiceConfig, clientConfig *email.Config) *dockertest.Resource {
	resource, err := pool.Run("rabbitmq", "3.11.0", []string{"RABBITMQ_DEFAULT_USER=root", "RABBITMQ_DEFAULT_PASS=root"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = resource.Expire(60)
	if err != nil {
		log.Fatalf("Could not set expire: %s", err)
	}

	url := fmt.Sprintf("amqp://root:root@localhost:%v/", resource.GetPort("5672/tcp"))
	serverConfig.Rabbitmq.RabbitMQConfig.RabbitUrl = url
	clientConfig.RabbitMQ.RabbitUrl = url

	return resource
}

func runMailHog(pool *dockertest.Pool, serverConfig *config.EmailServiceConfig) (*dockertest.Resource, string) {
	resource, err := pool.Run("mailhog/mailhog", "v1.0.1", []string{"TZ=Europe/Warsaw"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = resource.Expire(60)
	if err != nil {
		log.Fatalf("Could not set expire: %s", err)
	}

	serverConfig.Smtp.Port = resource.GetPort("1025/tcp")

	return resource, fmt.Sprintf("http://localhost:%v/", resource.GetPort("8025/tcp"))
}

func getMessages(t *testing.T, mailHogUrl string, startTime *time.Time) *mailhog.Messages {
	for i := 0; i < 10; i++ {
		messages, err := mailhog.GetMessages(mailHogUrl)
		assert.Nil(t, err)
		for _, message := range messages {
			if message.Created.After(*startTime) {
				return &message
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func runMockServer(pool *dockertest.Pool, config *conversion.Config) *dockertest.Resource {
	resource, err := pool.Run("mockserver/mockserver", "5.14.0", []string{})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	err = resource.Expire(60)
	if err != nil {
		log.Fatalf("Could not set expire: %s", err)
	}

	config.Url = fmt.Sprintf("http://localhost:%v", resource.GetPort("1080/tcp"))

	return resource
}

func waitForMockServer(config *conversion.Config) bool {
	for i := 0; i < 20; i++ {
		if httpGet(config) == http.StatusNotFound {
			return true
		}
		time.Sleep(100 * time.Millisecond)
	}
	return false
}

func httpGet(config *conversion.Config) int {
	httpClient := &http.Client{Timeout: 1 * time.Second}
	req, err := http.NewRequest("GET", config.Url, nil)
	if err != nil {
		return 0
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0
	}
	err = resp.Body.Close()
	if err != nil {
		return 0
	}
	return resp.StatusCode
}

func newTime(t *testing.T) *time.Time {
	location, err := time.LoadLocation("Europe/Warsaw")
	assert.Nil(t, err)
	resultTime := time.Now().In(location)
	return &resultTime
}

func newMockClient(t *testing.T, serverConfig *config.EmailServiceConfig) *mockclient.Client {
	return &mockclient.Client{
		T:       t,
		BaseURL: serverConfig.Conversion.Url,
	}
}

func CreateAttachments(fileName string, data string) map[string]*interface{} {
	attachments := make(map[string]*interface{})
	var dataInterface interface{}
	dataInterface = data
	attachments[fileName] = &dataInterface
	return attachments
}

func send(t *testing.T, client email.Client, email *rabbitMqModel.Email) {
	err := client.Send(email)
	assert.Nil(t, err)
}
