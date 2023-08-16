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
