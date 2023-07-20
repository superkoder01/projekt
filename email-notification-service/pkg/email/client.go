package email

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
	"go.uber.org/zap"
)

type Client interface {
	Send(email *model.Email) error
	Close() error
}

type client struct {
	producer messagebroker.MessageProducer
	logger   *zap.SugaredLogger
	config   *Config
}

func New(logger *zap.SugaredLogger, config *Config) (*client, error) {
	producer := rabbitmq.NewRabbitMQProducer(config.RabbitMQ, logger)
	err := producer.InitializeConnection()
	if err != nil {
		return nil, err
	}
	return &client{producer: producer, logger: logger, config: config}, nil
}

func (c *client) Send(email *model.Email) error {
	c.logger.Debugf("Send %v", email)

	bytes, err := ToBytes(email)
	if err != nil {
		c.logger.Errorf("Cannot encode mail. email: %v, err: %v", email, err)
	}

	return c.producer.PublishMessage(model.Message{Payload: bytes}, c.config.BindingKey)
}

func (c *client) Close() error {
	c.logger.Info("Close")
	return c.producer.Close()
}
