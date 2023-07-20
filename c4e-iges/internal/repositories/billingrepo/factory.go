package billingrepo

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/billingrepo/mariadb"
)

type mariadbFactory struct {
	cfg    *config.AppConfig
	logger logger.Logger
}

func NewMariadbFactory(cfg *config.AppConfig, logger logger.Logger) *mariadbFactory {
	return &mariadbFactory{
		cfg:    cfg,
		logger: logger,
	}
}

func (s *mariadbFactory) MakeRepo() ports.IBillingRepo {
	return mariadb.NewBillingRepo(s.cfg, s.logger)
}
