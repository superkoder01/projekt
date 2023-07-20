package rabbit_subscriber

import (
	"context"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/worker"
	"strconv"
)

type rabbitSubscriber struct {
	context.Context
	messagebroker.MessageConsumer
	log logger.Logger
	cfg *config.AppConfig
}

func NewRabbitSubscriber(ctx context.Context, worker worker.Worker, log logger.Logger, cfg *config.AppConfig) *rabbitSubscriber {
	log.Info("initializing rabbitmq invoice event subscriber")

	consumer := rabbitmq.NewRabbitMQConsumer(cfg.RabbitConsumer, worker, log)
	if err := consumer.InitializeConnection(); err != nil {
		log.Errorf("failed to initialize rabbitmq invoice event subscriber, error: %v", err)
	}

	return &rabbitSubscriber{ctx, consumer, log, cfg}
}

func (r *rabbitSubscriber) Subscribe(ctx context.Context) chan error {
	consumerName := "InvoiceEventConsumer"
	chanErr := make(chan error)

	r.log.Infof("subscribing for invoice events in queue: %s", r.cfg.RabbitConsumer.QueueName)

	for i := 1; i <= r.cfg.InvoiceService.InvoiceEventConsumerPoolSize; i++ {
		go func(index string) {
			err := r.StartConsumer(fmt.Sprintf("%s-%s", consumerName, index))
			if err != nil {
				chanErr <- err
			}
		}(strconv.Itoa(i))
	}

	r.log.Infof("number of invoice event subscribers listening: (%d)", r.cfg.InvoiceService.InvoiceEventConsumerPoolSize)

	return chanErr
}

func (r *rabbitSubscriber) Close() error {
	r.log.Info("invoice event subscriber stopping")
	return r.MessageConsumer.Close()
}
