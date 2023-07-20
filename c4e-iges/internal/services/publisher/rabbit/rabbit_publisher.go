package rabbit_publisher

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
)

type rabbitPublisher struct {
	messagebroker.MessageProducer
	log logger.Logger
	cfg *config.AppConfig
}

func NewPublisher(log logger.Logger, cfg *config.AppConfig) *rabbitPublisher {
	producer := rabbitmq.NewRabbitMQProducer(cfg.RabbitProducer, log)
	err := producer.InitializeConnection()
	if err != nil {
		log.Fatalf("failed to initialize rabbitmq invoice publisher, error: %v", err)
	}

	return &rabbitPublisher{producer, log, cfg}
}

func (rp *rabbitPublisher) Publish(ctx context.Context, invoice *billing.InvoiceProsument) error {
	rp.log.Infof("publishing invoice, customerId: %s, number: %s",
		invoice.Payload.CustomerDetails.CustomerId,
		invoice.Payload.InvoiceDetails.Number)

	if message, err := CreateEmailMessage(
		[]string{invoice.Payload.CustomerDetails.Contact.Email},
		rp.cfg.InvoiceService.EmailTitleTemplate,
		rp.cfg.InvoiceService.EmailBodyTemplate,
		*invoice,
	); err != nil {
		return err
	} else {
		//rp.log.Debugf("email: %v", message)
		if err = rp.PublishMessage(message, "email"); err != nil {
			return err
		}
	}
	return nil
}

func (rp *rabbitPublisher) Close() error {
	rp.log.Info("invoice publisher stopping")
	return rp.MessageProducer.Close()
}
