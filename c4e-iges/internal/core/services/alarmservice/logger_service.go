package alarmservice

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain/alarms"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
)

type loggerAalarmService struct {
	log logger.Logger
	cfg *config.AppConfig
}

func newAlarmService(log logger.Logger, cfg *config.AppConfig) ports.AlarmService {
	return &loggerAalarmService{log: log, cfg: cfg}
}

func (s loggerAalarmService) SendAlarm(ctx context.Context, alarm *alarms.Alarm) {
	s.log.Error(alarm.String())
}
