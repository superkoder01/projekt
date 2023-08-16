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
package eventhandler

import (
	"encoding/json"
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/scheduler"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing.git/pkg/invoice"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/actors"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/validators"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/factory"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
)

type actorSystem struct {
	system      *actor.ActorSystem
	decider     func(reason interface{}) actor.Directive
	supervisor  actor.SupervisorStrategy
	rootContext *actor.RootContext
	actorProps  *actor.Props
	scheduler   *scheduler.TimerScheduler
}

type eventHandler struct {
	actorSystem
	log         logger.Logger
	cfg         *config.AppConfig
	faktory     factory.Factory
	initialized bool
}

func NewEventHandler(faktory factory.Factory, log logger.Logger, cfg *config.AppConfig) *eventHandler {
	return &eventHandler{
		log:     log,
		cfg:     cfg,
		faktory: faktory,
	}
}

func (eh *eventHandler) Initialize() error {
	if eh.initialized {
		eh.log.Warn("service already initialized")
		return nil
	}

	eh.system = actor.NewActorSystem()
	eh.decider = func(reason interface{}) actor.Directive {
		eh.log.Warn("handling failure for child")
		return actor.StopDirective
	}
	eh.supervisor = actor.NewOneForOneStrategy(10, 1000, eh.decider)
	eh.rootContext = eh.system.Root
	eh.scheduler = scheduler.NewTimerScheduler(eh.rootContext)
	eh.actorProps = actor.
		PropsFromProducer(
			actors.NewProsument(
				&struct {
					ports.ContractRepo
					ports.InvoiceRepo
				}{eh.faktory.ContractRepo(), eh.faktory.InvoiceRepo()},
				&struct {
					ports.InvoicePublisher
					ports.AlarmService
				}{eh.faktory.InvoicePublisher(), eh.faktory.AlarmService()},
				&struct {
					ports.ContractValidator
					ports.InvoiceEventValidator
				}{validators.NewContractValidator(eh.log, eh.cfg), validators.NewInvoiceEventValidator(eh.log, eh.cfg)},
				eh.log, eh.cfg),
			actor.WithSupervisor(eh.supervisor))

	eh.initialized = true
	eh.log.Info("event handler initialized")
	return nil
}

func (eh *eventHandler) Execute(message []byte, messageId, consumerName string) error {
	var event *invoice.InvoiceEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		eh.log.Errorf("failed to deserialize event: %v", err)
		return err
	}

	eh.log.Debugf("%s consume event with id: %v", consumerName, messageId)

	pid, err := eh.rootContext.SpawnNamed(eh.actorProps, fmt.Sprintf("PROSUMENT-%s", event.Contract))
	if err != nil {
		eh.log.Errorf("failed to create actor: %v", err)
		return err
	}

	eh.scheduler.SendOnce(0, pid, &actors.BillingMessage{InvoiceEvent: event})
	//eh.rootContext.Send(pid, &actors.BillingMessage{InvoiceEvent: event})

	return nil
}
