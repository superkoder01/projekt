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
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain/alarms"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/generators"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	"time"
)

// TODO implement support for B2B
// TODO do we need separate child actors for generating invoices or it can be done in parent actor?
// TODO implement alarm generation
// TODO review and fix logging to be more descriptive (fix log levels where appropriate)
// TODO test against failure scenarios, deal with timeouts, errors, etc. fix context usage where appropriate

type (
	ProsumentType int

	contractFetched struct {
		contract *billing.Contract
	}
	fetchError struct {
		Error error
	}

	invoiceEventValidated     struct{}
	invoiceEventValidateError struct {
		Error error
	}

	contractValidated     struct{}
	contractValidateError struct {
		Error error
	}

	invoiceStored                  struct{}
	invoiceDetailsStored           struct{}
	repurchaseInvoiceStored        struct{}
	repurchaseInvoiceDetailsStored struct{}
	storeError                     struct {
		Error error
	}

	settlementPublished    struct{}
	settlementPublishError struct {
		Error error
	}

	invoiceNumberGenerated struct {
		number string
	}
	invoiceNumberGenerateError struct {
		Error error
	}

	repurchaseInvoiceNumberGenerated struct {
		number string
	}
	repurchaseInvoiceNumberGenerateError struct {
		Error error
	}

	validator interface {
		ports.ContractValidator
		ports.InvoiceEventValidator
	}

	repository interface {
		ports.ContractRepo
		ports.InvoiceRepo
	}

	publisher interface {
		ports.InvoicePublisher
		ports.AlarmService
	}

	billingDetails struct {
		contract     *billing.Contract
		invoiceEvent *invoice.InvoiceEvent
	}

	settlement struct {
		invoiceNumber           string
		repurchaseInvoiceNumber string
		documents               []interface{}
	}

	numberGenerator struct {
		sellInvoiceNumberGenerator       ports.InvoiceNumberGenerator
		repurchaseInvoiceNumberGenerator ports.InvoiceNumberGenerator
	}

	base struct {
		behavior actor.Behavior
		log      logger.Logger
		cfg      *config.AppConfig
		repository
		publisher
		validator
		billingDetails
		settlement
		numberGenerator
	}

	// Prosument Actor
	prosument struct {
		base
		prosumentType ProsumentType
	}
)

const (
	B2C ProsumentType = iota
	B2B
)

func NewProsument(repository repository, publisher publisher, validator validator, log logger.Logger, cfg *config.AppConfig) actor.Producer {
	return func() actor.Actor {
		p := &prosument{
			base: base{
				log:            log,
				cfg:            cfg,
				repository:     repository,
				publisher:      publisher,
				validator:      validator,
				billingDetails: billingDetails{},
				settlement:     settlement{},
				numberGenerator: numberGenerator{
					sellInvoiceNumberGenerator:       generators.NewInvoiceNumberGenerator(repository, log, cfg),
					repurchaseInvoiceNumberGenerator: generators.NewRepurchaseInvoiceNumberGenerator(repository, log, cfg),
				},
			},
			prosumentType: B2C,
		}
		p.log.Debugf("actor created")
		p.behavior.Become(p.initialState())
		return p
	}
}

func NewProsumentB2B(repository repository, publisher publisher, validator validator, log logger.Logger, cfg *config.AppConfig) actor.Producer {
	return func() actor.Actor {
		p := &prosument{
			base: base{
				log:            log,
				cfg:            cfg,
				repository:     repository,
				publisher:      publisher,
				validator:      validator,
				billingDetails: billingDetails{},
				settlement:     settlement{},
				numberGenerator: numberGenerator{
					sellInvoiceNumberGenerator:       generators.NewInvoiceNumberGenerator(repository, log, cfg),
					repurchaseInvoiceNumberGenerator: generators.NewRepurchaseInvoiceNumberGenerator(repository, log, cfg),
				},
			},
			prosumentType: B2B,
		}
		p.log.Debugf("actor created")
		p.behavior.Become(p.initialState())
		return p
	}
}

func (a *prosument) trace(template string, args ...interface{}) string {
	msg := fmt.Sprintf(template, args...)
	a.log.Debugf("entering: " + msg)
	return msg
}

func (a *prosument) un(msg string) {
	a.log.Debug("leaving: " + msg)
}

func (a *prosument) Receive(actx actor.Context) {
	defer a.un(a.trace("[%v] received %T", actx.Self(), actx.Message()))
	a.behavior.Receive(actx)
}

func (a *prosument) initialState() actor.ReceiveFunc {
	defer a.un(a.trace("%s", "initialState"))

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *BillingMessage:
			a.invoiceEvent = m.InvoiceEvent
			a.behavior.Become(a.validateInvoiceEventState(c, nil))
		}
	}
}

func (a *prosument) validateInvoiceEventState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] validateInvoiceEventState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		go a.validateInvoiceEvent(ctx, actx)
		responses = 1
	}

	if responses == 0 {
		cancel()
		return a.fetchContractState(actx, nil)
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *invoiceEventValidated:
			a.behavior.Become(a.validateInvoiceEventState(c, cancel, responses-1))
		case *invoiceEventValidateError:
			msg := fmt.Sprintf("[%v] invoice event validation failed, reason: %v", c.Self(), m.Error)
			a.log.Errorf("%v", msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) fetchContractState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] fetchContractState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		go a.fetchContract(ctx, actx)
		responses = 1
	}

	if responses == 0 {
		cancel()
		return a.validateContractState(actx, nil)
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *contractFetched:
			a.contract = m.contract
			a.behavior.Become(a.fetchContractState(c, cancel, responses-1))
		case *fetchError:
			msg := fmt.Sprintf("[%v] could not fetch data, reason: %v", c.Self(), m.Error)
			a.log.Errorf("%v", msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) validateContractState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] contractValidationState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		go a.validateContract(ctx, actx)
		responses = 1
	}

	if responses == 0 {
		cancel()
		return a.generateInvoiceNumberState(actx, nil)
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *contractValidated:
			a.behavior.Become(a.validateContractState(c, cancel, responses-1))
		case *contractValidateError:
			msg := fmt.Sprintf("[%v] contract validation failed, reason: %v", c.Self(), m.Error)
			a.log.Errorf("%v", msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) generateInvoiceNumberState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] generateInvoiceNumberState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		go a.generateInvoiceNumber(ctx, actx)
		go a.generateRepurchaseInvoiceNumber(ctx, actx)
		responses = 2
	}

	if responses == 0 {
		cancel()
		return a.generateSettlementState(actx, nil)
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *invoiceNumberGenerated:
			a.invoiceNumber = m.number
			a.behavior.Become(a.generateInvoiceNumberState(c, cancel, responses-1))
		case *repurchaseInvoiceNumberGenerated:
			a.repurchaseInvoiceNumber = m.number
			a.behavior.Become(a.generateInvoiceNumberState(c, cancel, responses-1))
		case *invoiceNumberGenerateError:
			msg := fmt.Sprintf("[%v] invoice number generation failed, reason: %v", c.Self(), m.Error)
			a.log.Errorf("%v", msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		case *repurchaseInvoiceNumberGenerateError:
			msg := fmt.Sprintf("[%v] repurchase invoice number generation failed, reason: %v", c.Self(), m.Error)
			a.log.Errorf("%v", msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) generateSettlementState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] generateSettlementState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		responses = a.spawnGenerators(ctx, actx)
	}

	if responses == 0 {
		cancel()
		return a.storeSettlementState(actx, nil)
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *InvoiceGenerateResponse:
			a.documents = append(a.documents, m.invoice)
			a.documents = append(a.documents, m.invoiceDetails)
			a.behavior.Become(a.generateSettlementState(c, cancel, responses-1))
		case *RepurchaseInvoiceGenerateResponse:
			a.documents = append(a.documents, m.repurchaseInvoice)
			a.documents = append(a.documents, m.repurchaseInvoiceDetails)
			a.behavior.Become(a.generateSettlementState(c, cancel, responses-1))
		case *InvoiceGenerateError:
			msg := fmt.Sprintf("[%v] invoice generation failed, reason: %s", c.Self(), m.Error.Error())
			a.log.Errorf(msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		case *RepurchaseInvoiceGenerateError:
			msg := fmt.Sprintf("[%v] repurchase invoice generation failed, reason: %s", c.Self(), m.Error.Error())
			a.log.Errorf(msg)
			cancel()
			a.generateAlarm(fmt.Errorf(msg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) storeSettlementState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] storeSettlementState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		go a.storeInvoices(ctx, actx)
		responses = 1
	}

	if responses == 0 {
		cancel()
		return a.publishSettlementState(actx, nil)
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *invoiceStored:
			a.behavior.Become(a.storeSettlementState(c, cancel, responses-1))
		case *invoiceDetailsStored:
			a.behavior.Become(a.storeSettlementState(c, cancel, responses-1))
		case *repurchaseInvoiceStored:
			a.behavior.Become(a.storeSettlementState(c, cancel, responses-1))
		case *repurchaseInvoiceDetailsStored:
			a.behavior.Become(a.storeSettlementState(c, cancel, responses-1))
		case *storeError:
			msg := fmt.Sprintf("[%v] could not store data, reason: %s", c.Self(), m.Error)
			a.log.Errorf(msg)
			cancel()
			a.generateAlarm(fmt.Errorf("%v", msg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) publishSettlementState(actx actor.Context, cancel context.CancelFunc, expectedResponse ...int) actor.ReceiveFunc {
	responses := getFirstOrZero(expectedResponse...)
	defer a.un(a.trace("[%v] publishSettlementState (%d)", actx.Self(), responses))

	if cancel == nil {
		ctx := context.Background()
		ctx, cancel = context.WithTimeout(ctx, time.Second*300)
		//defer cancel()
		go a.publishSettlement(ctx, actx)
		responses = 1
	}

	if responses == 0 {
		cancel()
		actx.Stop(actx.Self())
	}

	return func(c actor.Context) {
		switch m := c.Message().(type) {
		case *settlementPublished:
			a.behavior.Become(a.publishSettlementState(c, cancel, responses-1))
		case *settlementPublishError:
			errMsg := fmt.Sprintf("[%v] unable to publish invoice, reason: %v", c.Self(), m.Error)
			a.log.Errorf(errMsg)
			cancel()
			a.generateAlarm(fmt.Errorf("%v", errMsg), c)
			c.Stop(c.Self())
		}
	}
}

func (a *prosument) fetchContract(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "fetchContract"))

	var msg interface{}

	contract, err := a.repository.GetContractByContractNumber(ctx, a.invoiceEvent.Contract)
	if err != nil {
		msg = &fetchError{
			Error: err,
		}
	} else {
		msg = &contractFetched{
			contract: contract,
		}
	}

	actx.Send(actx.Self(), msg)
}

func (a *prosument) generateAlarm(err error, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "generateAlarm"))

	a.SendAlarm(context.Background(), alarms.OnGeneralError(err))
}

func (a *prosument) spawnGenerators(ctx context.Context, actx actor.Context) int {
	defer a.un(a.trace("[%v] %s", actx.Self(), "spawnGenerators"))

	responses := 0
	switch a.prosumentType {
	case B2C:
		responses = 2
		a.spawnInvoiceGenerator(ctx, actx, generators.NewInvoiceGeneratorB2C(a.contract, a.invoiceEvent, a.log, a.cfg))
		a.spawnRepurchaseInvoiceGenerator(ctx, actx, generators.NewRepurchaseInvoiceGeneratorB2C(a.contract, a.invoiceEvent, a.log, a.cfg))
	case B2B:
		responses = 1
		a.spawnInvoiceGenerator(ctx, actx, generators.NewInvoiceGeneratorB2B(a.contract, a.invoiceEvent, a.log, a.cfg))
		//a.spawnRepurchaseInvoiceGenerator(ctx, actx, services.NewRepurchaseInvoiceGeneratorB2B(a.contract, a.invoiceEvent, a.log, a.cfg))
	}

	return responses
}

func (a *prosument) spawnInvoiceGenerator(ctx context.Context, parentActx actor.Context, service ports.InvoiceGenerator) {
	defer a.un(a.trace("[%v] spawning invoice generation, contract: %s, customerId: %s", parentActx.Self(), a.invoiceEvent.Contract, a.contract.Payload.ContractDetails.CustomerId))

	props := actor.PropsFromProducer(newInvoiceGenerator(service, a.log, a.cfg))
	child, _ := parentActx.SpawnNamed(props, "INVOICE")
	parentActx.Send(child, &InvoiceGenerateRequest{a.invoiceNumber})
}

func (a *prosument) spawnRepurchaseInvoiceGenerator(ctx context.Context, parentActx actor.Context, service ports.RepurchaseInvoiceGenerator) {
	defer a.un(a.trace("[%v] spawning repurchase invoice generation, contract: %s, customerId: %s", parentActx.Self(), a.invoiceEvent.Contract, a.contract.Payload.ContractDetails.CustomerId))

	props := actor.PropsFromProducer(newRepurchaseInvoiceGenerator(service, a.log, a.cfg))
	child, _ := parentActx.SpawnNamed(props, "REPURCHASE-INVOICE")
	parentActx.Send(child, &RepurchaseInvoiceGenerateRequest{a.repurchaseInvoiceNumber})
}

func (a *prosument) storeInvoices(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "storeInvoices"))

	var msg interface{} = &invoiceStored{}

	err := a.repository.StoreManyWithinTransaction(ctx, a.documents...)
	if err != nil {
		msg = storeError{
			Error: err,
		}
	}

	actx.Send(actx.Self(), msg)
}

func (a *prosument) validateInvoiceEvent(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "validateInvoiceEvent"))

	var msg interface{} = &invoiceEventValidated{}

	if err := a.ValidateInvoiceEvent(ctx, a.invoiceEvent); err != nil {
		msg = &invoiceEventValidateError{Error: err}
	}

	actx.Send(actx.Self(), msg)
}

func (a *prosument) validateContract(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "validateContract"))

	var msg interface{} = &contractValidated{}

	if err := a.ValidateContract(ctx, a.invoiceEvent, a.contract); err != nil {
		msg = &contractValidateError{Error: err}
	}

	actx.Send(actx.Self(), msg)
}

func (a *prosument) publishSettlement(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "publishSettlement"))

	var msg interface{} = &settlementPublished{}

	// TODO currently sending of InvoiceProsument is supported
	for _, document := range a.documents {
		switch doc := document.(type) {
		case *billing.InvoiceProsument:
			a.log.Debugf("publishing %T", doc)
			if err := a.Publish(ctx, doc); err != nil {
				msg = &settlementPublishError{Error: err}
			}
		default:
			a.log.Warnf("trying to publish unknown document type %T - skip publishing", doc)
		}
	}

	actx.Send(actx.Self(), msg)
}

func (a *prosument) generateInvoiceNumber(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "generateInvoiceNumber"))

	var msg interface{}

	number, err := a.sellInvoiceNumberGenerator.GetNumber(ctx, a.contract.Payload.CustomerDetails.CustomerId, *a.invoiceEvent)
	if err != nil {
		msg = &invoiceNumberGenerateError{
			Error: err,
		}
	} else {
		msg = &invoiceNumberGenerated{
			number: number,
		}
	}

	actx.Send(actx.Self(), msg)
}

func (a *prosument) generateRepurchaseInvoiceNumber(ctx context.Context, actx actor.Context) {
	defer a.un(a.trace("[%v] %s", actx.Self(), "generateRepurchaseInvoiceNumber"))

	var msg interface{}

	number, err := a.repurchaseInvoiceNumberGenerator.GetNumber(ctx, a.contract.Payload.CustomerDetails.CustomerId, *a.invoiceEvent)
	if err != nil {
		msg = &repurchaseInvoiceNumberGenerateError{
			Error: err,
		}
	} else {
		msg = &repurchaseInvoiceNumberGenerated{
			number: number,
		}
	}

	actx.Send(actx.Self(), msg)
}

func getFirstOrZero(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	return values[0]
}
