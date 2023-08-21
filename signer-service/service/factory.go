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
