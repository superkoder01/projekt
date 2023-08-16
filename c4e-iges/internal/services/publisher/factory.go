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
package publisher

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	dummy_publisher "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/services/publisher/dummy"
	rabbit_publisher "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/services/publisher/rabbit"
)

type PublisherType string

const (
	RABBITMQ PublisherType = "rabbitmq"
	DUMMY                  = "dummy"
)

// publisher factory
type publisherFactory struct {
	context.Context
	publisherType PublisherType
	log           logger.Logger
	cfg           *config.AppConfig
}

func NewPublisherFactory(ctx context.Context, publisherType PublisherType, log logger.Logger, cfg *config.AppConfig) *publisherFactory {
	return &publisherFactory{ctx, publisherType, log, cfg}
}

func (f *publisherFactory) MakePublisher() ports.InvoicePublisher {
	switch f.publisherType {
	case RABBITMQ:
		return rabbit_publisher.NewPublisher(f.log, f.cfg)
	case DUMMY:
		return dummy_publisher.NewPublisher(f.log, f.cfg)
	default:
		panic("unknown publisher type: " + f.publisherType)
	}
}
