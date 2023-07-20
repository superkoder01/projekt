package alarmservice

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/ports"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
)

// AlarmServiceType alarm service type
type AlarmServiceType string

const (
	LOGGER AlarmServiceType = "logger"
)

//alarmServiceFactory
type alarmServiceFactory struct {
	context.Context
	alarmServiceType AlarmServiceType
	log              logger.Logger
	cfg              *config.AppConfig
}

//NewAlarmServiceFactory create alarm service factory
func NewAlarmServiceFactory(ctx context.Context, alarmServiceType AlarmServiceType, log logger.Logger, cfg *config.AppConfig) *alarmServiceFactory {
	return &alarmServiceFactory{ctx, alarmServiceType, log, cfg}
}

//MakeService create service
func (f *alarmServiceFactory) MakeService() ports.AlarmService {
	switch f.alarmServiceType {
	case LOGGER:
		return newAlarmService(f.log, f.cfg)
	default:
		panic("unknown service: " + f.alarmServiceType)
	}
}
