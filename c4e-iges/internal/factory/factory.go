package factory

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
)

type Factory interface {
	ContractRepo() ports.ContractRepo
	InvoiceRepo() ports.InvoiceRepo
	InvoicePublisher() ports.InvoicePublisher
	AlarmService() ports.AlarmService
}

type factory struct {
	contractRepo     ports.ContractRepo
	invoiceRepo      ports.InvoiceRepo
	invoicePublisher ports.InvoicePublisher
	alarmService     ports.AlarmService
}

func (f *factory) ContractRepo() ports.ContractRepo {
	return f.contractRepo
}

func (f *factory) InvoiceRepo() ports.InvoiceRepo {
	return f.invoiceRepo
}

func (f *factory) InvoicePublisher() ports.InvoicePublisher {
	return f.invoicePublisher
}

func (f *factory) AlarmService() ports.AlarmService {
	return f.alarmService
}

// builder
type factoryMod func(f *factory)
type factoryBuilder struct {
	actions []factoryMod
}

func NewFactoryBuilder() *factoryBuilder {
	return &factoryBuilder{}
}

func (f *factoryBuilder) WithContractRepo(contractRepo ports.ContractRepo) *factoryBuilder {
	f.actions = append(f.actions, func(p *factory) {
		p.contractRepo = contractRepo
	})
	return f
}

func (f *factoryBuilder) WithInvoiceRepo(invoiceRepo ports.InvoiceRepo) *factoryBuilder {
	f.actions = append(f.actions, func(p *factory) {
		p.invoiceRepo = invoiceRepo
	})
	return f
}

func (f *factoryBuilder) WithInvoicePublisher(invoicePublisher ports.InvoicePublisher) *factoryBuilder {
	f.actions = append(f.actions, func(p *factory) {
		p.invoicePublisher = invoicePublisher
	})
	return f
}

func (f *factoryBuilder) WithAlarmService(alarmService ports.AlarmService) *factoryBuilder {
	f.actions = append(f.actions, func(p *factory) {
		p.alarmService = alarmService
	})
	return f
}

func (f *factoryBuilder) Create() Factory {
	p := &factory{}
	for _, action := range f.actions {
		action(p)
	}
	return p
}
