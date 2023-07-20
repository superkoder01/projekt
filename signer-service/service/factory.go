package service

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/service/impl"
	"go.uber.org/zap"
)

type ServiceFactory interface {
	New(string) Service
}

type serviceFactory struct {
	logger *zap.SugaredLogger
}

func NewServiceFactory(logger *zap.SugaredLogger) *serviceFactory {
	return &serviceFactory{logger: logger}
}

const (
	SIGNER = "SIGNER"
)

func (sf *serviceFactory) New(name string) Service {
	switch name {
	case SIGNER:
		conversionConfig := conversion.Config{Url: configuration.GetConversionServiceConfigConfig().Url, Timeout: configuration.GetConversionServiceConfigConfig().Timeout}
		return impl.NewReportService(sf.logger, conversion.NewByConfig(&conversionConfig, sf.logger), configuration.GetMSzafirConfig())
	default:
		return nil

	}

}
