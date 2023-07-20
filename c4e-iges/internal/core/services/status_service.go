package services

import (
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
)

type status struct {
	log logger.Logger
	cfg *config.AppConfig
}

func NewStatusService(log logger.Logger, cfg *config.AppConfig) ports.Status {
	return &status{
		log: log,
		cfg: cfg,
	}
}

func (s *status) IsAlive() (bool, error) {
	s.log.Info("service is alive")
	return true, nil
}
