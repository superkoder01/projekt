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
	repurchaseInvoiceGeneratorActor struct {
		log     logger.Logger
		cfg     *config.AppConfig
		service ports.RepurchaseInvoiceGenerator
	}
)

func newRepurchaseInvoiceGenerator(service ports.RepurchaseInvoiceGenerator, log logger.Logger, cfg *config.AppConfig) actor.Producer {
	return func() actor.Actor {
		return &repurchaseInvoiceGeneratorActor{
			log:     log,
			cfg:     cfg,
			service: service,
		}
	}
}

func (a *repurchaseInvoiceGeneratorActor) trace(template string, args ...interface{}) string {
	msg := fmt.Sprintf(template, args...)
	a.log.Debugf("entering: " + msg)
	return msg
}

func (a *repurchaseInvoiceGeneratorActor) un(msg string) {
	a.log.Debug("leaving: " + msg)
}

func (a *repurchaseInvoiceGeneratorActor) Receive(actx actor.Context) {
	defer a.un(a.trace("[%v] received %T", actx.Self(), actx.Message()))

	switch msg := actx.Message().(type) {
	case *RepurchaseInvoiceGenerateRequest:
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, time.Second*300)
		defer cancel()
		repurchaseInvoice, repurchaseInvoiceDetails, err := a.service.GenerateRepurchaseInvoice(ctx, msg.number)

		if err != nil {
			actx.Send(actx.Parent(), &RepurchaseInvoiceGenerateError{err})
		} else {
			response := &RepurchaseInvoiceGenerateResponse{repurchaseInvoice: repurchaseInvoice, repurchaseInvoiceDetails: repurchaseInvoiceDetails}
			actx.Send(actx.Parent(), response)
		}
	}
}
