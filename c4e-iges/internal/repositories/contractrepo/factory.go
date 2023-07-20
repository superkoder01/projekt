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
