package ports

import (
	"context"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/core/domain/alarms"
)

type (
	AlarmService interface {
		SendAlarm(ctx context.Context, alarm *alarms.Alarm)
	}

	AlarmServiceFactory interface {
		MakeService() AlarmService
	}
)
