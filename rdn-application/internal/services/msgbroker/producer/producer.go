package producer

import (
	"RDN-application/internal/ports"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
)

type rabbitProducer struct {
	producer messagebroker.MessageProducer
	config   config.AppConfig
	logger   logger.Logger
}

func NewNotificationService(cfg config.AppConfig, logger logger.Logger) *rabbitProducer {
	rabbitProducer := &rabbitProducer{
		producer: rabbitmq.NewRabbitMQProducer(*cfg.GetRabbitMqProducerConfig(), logger),
		config:   cfg,
		logger:   logger,
	}
	err := rabbitProducer.producer.InitializeConnection()
	if err != nil {
		logger.Fatalf("Failed to init connection with RabbitMq. %v", err)
	}
	return rabbitProducer
}

func (rabbit *rabbitProducer) PublishToUser(ctx context.Context, messageBody string, userAddress []string, notificationType ports.NotificationType) error {
	rabbit.logger.Infof("Publishing message(%v) to: %v", notificationType, userAddress)
	/*todo use context*/

	message, err := rabbit.createMessageWithSpecificType(notificationType, userAddress, messageBody)
	if err != nil {
		rabbit.logger.Errorf("Failed to create message - Cannot publish !")
		return err
	} else {
		return rabbit.producer.PublishMessage(message, string(notificationType))
	}
}

func (rabbit *rabbitProducer) createMessageWithSpecificType(
	notificationType ports.NotificationType, userAddress []string, messageBody string) (model.Message, error) {
	var message interface{}
	var be bytes.Buffer
	encoder := json.NewEncoder(&be)

	switch notificationType {
	case ports.Email:
		message = model.Email{
			Destination: userAddress,
			Title:       rabbit.config.GetServiceConfig().ServiceName,
			Body:        "Failed to fetch data\nError: " + messageBody + "\nPlease check logs and service configuration",
		}
	case ports.Sms:
		message = model.Sms{
			Sender: "0048660555888",
			Msisdn: userAddress,
			Text:   "[" + rabbit.config.GetServiceConfig().ServiceName + " Error] Failed to fetch data: " + messageBody + ". Please check logs",
		}
	default:
		rabbit.logger.Errorf("Invalid notification type !!", notificationType)
		return model.Message{}, errors.New("invalid notification type")
	}

	err := encoder.Encode(message)
	if err != nil {
		return model.Message{}, err
	}

	return model.Message{Payload: be.Bytes()}, nil
}
