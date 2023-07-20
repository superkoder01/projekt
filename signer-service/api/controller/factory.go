package controller

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/api/controller/impl"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/service"
)

type ControllerFactory interface {
	New(string) Controller
}

type controllerFactory struct {
	sf service.ServiceFactory
}

func NewControllerFactory(s service.ServiceFactory) *controllerFactory {
	return &controllerFactory{sf: s}
}

const (
	SIGNER = "SIGNER"
)

func (hf *controllerFactory) New(name string) Controller {
	switch name {
	case SIGNER:
		return impl.NewReportController(hf.sf.New(SIGNER))
	default:
		return nil
	}
}
