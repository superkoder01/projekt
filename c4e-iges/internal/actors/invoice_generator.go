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
