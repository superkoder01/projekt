package actors

import (
	"context"
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"time"
)

type (
	invoiceGeneratorActor struct {
		log     logger.Logger
		cfg     *config.AppConfig
		service ports.InvoiceGenerator
	}
)

func newInvoiceGenerator(service ports.InvoiceGenerator, log logger.Logger, cfg *config.AppConfig) actor.Producer {
	return func() actor.Actor {
		return &invoiceGeneratorActor{
			log:     log,
			cfg:     cfg,
			service: service,
		}
	}
}

func (a *invoiceGeneratorActor) trace(template string, args ...interface{}) string {
	msg := fmt.Sprintf(template, args...)
	a.log.Debugf("entering: " + msg)
	return msg
}

func (a *invoiceGeneratorActor) un(msg string) {
	a.log.Debug("leaving: " + msg)
}

func (a *invoiceGeneratorActor) Receive(actx actor.Context) {
	defer a.un(a.trace("[%v] received %T", actx.Self(), actx.Message()))

	switch msg := actx.Message().(type) {
	case *InvoiceGenerateRequest:
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second*300)
		defer cancel()
		invoice, invoiceDetails, err := a.service.GenerateInvoice(ctx, msg.number)

		if err != nil {
			actx.Send(actx.Parent(), &InvoiceGenerateError{err})
		} else {
			response := &InvoiceGenerateResponse{invoice: invoice, invoiceDetails: invoiceDetails}
			actx.Send(actx.Parent(), response)
		}
	}
}
