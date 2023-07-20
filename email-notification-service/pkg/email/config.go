package email

import "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"

type Config struct {
	BindingKey string
	RabbitMQ   config.RabbitMQProducerConfig
}
