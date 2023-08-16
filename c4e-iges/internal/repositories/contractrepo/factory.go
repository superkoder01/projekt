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
package contractrepo

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/contractrepo/mongo"
)

type ContractRepoType string

const (
	MONGO ContractRepoType = "mongo"
)

type contractRepoFactory struct {
	context.Context
	contractRepoType ContractRepoType
	log              logger.Logger
	cfg              *config.AppConfig
}

func NewContractRepoFactory(ctx context.Context, contractRepoType ContractRepoType, log logger.Logger, cfg *config.AppConfig) *contractRepoFactory {
	return &contractRepoFactory{ctx, contractRepoType, log, cfg}
}

func (f *contractRepoFactory) MakeRepo() ports.ContractRepo {
	switch f.contractRepoType {
	case MONGO:
		return mongo.NewContractRepo(f.log, f.cfg)
	default:
		panic("unknown repo: " + f.contractRepoType)
	}
}
