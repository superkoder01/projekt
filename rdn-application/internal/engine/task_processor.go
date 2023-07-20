package engine

import (
	"RDN-application/internal/model"
	"RDN-application/internal/ports"
	"RDN-application/pkg/config"
	"RDN-application/pkg/logger"
	"context"
	"errors"
	"time"
)

type taskProcessor struct {
	collector       ports.DataCollector
	dataService     ports.DataService
	notifier        ports.NotificationRepo
	numberOfRetries int
	retryDelay      time.Duration
	config          config.AppConfig
	logger          logger.Logger
}

func NewTaskProcessor(collector ports.DataCollector, dataService ports.DataService, notifier ports.NotificationRepo,
	numberOfRetries int, appConfig config.AppConfig, logger logger.Logger) *taskProcessor {
	return &taskProcessor{
		config:          appConfig,
		collector:       collector,
		dataService:     dataService,
		notifier:        notifier,
		numberOfRetries: numberOfRetries,
		logger:          logger,
		retryDelay:      time.Duration(appConfig.GetCollectorConfig().RetryTime),
	}
}

func (processor *taskProcessor) ProcessCollectionTask() func(ctx context.Context, errChan chan error, jobDate []time.Time) {
	return func(ctx context.Context, errChan chan error, jobDate []time.Time) {
		date := time.Now().Local()
		if jobDate != nil && len(jobDate) >= 1 {
			date = jobDate[0]
		}

		var dataChan = make(chan model.Collection)
		collectCtx, shutdown := context.WithTimeout(ctx, 60*time.Second)
		go processor.collector.CollectDataFromDate(collectCtx, date, dataChan)

		select {
		case <-ctx.Done():
			shutdown()
			errChan <- errors.New("context closed")
		case data := <-dataChan:
			shutdown()
			collectedData, collectDate := data.GetCollectionData()
			if data != nil && collectedData != nil && len(collectedData) > 1 && collectDate.IsZero() == false {
				err := processor.dataService.SaveDataToRepo(ctx, data, collectDate)
				if err != nil {
					processor.logger.Errorf("Failed to save data to database !")
					errChan <- err
				} else {
					processor.logger.Infof("Task finished")
				}
			} else {
				processor.logger.Errorf("No valid data from collector ! [%v %v]", collectedData, collectDate)
				errChan <- errors.New("empty dataset")
			}
		}
	}
}

func (processor *taskProcessor) WithFailureHandler(ctx context.Context, errChan chan error) {
	processingError := <-errChan
	processor.logger.Infof("Processing error: \"%v\"", processingError)

	for i := 0; i < processor.numberOfRetries; i++ {
		time.Sleep(processor.retryDelay * time.Second)
		processor.logger.Infof("Retrying...")

		retryChan := make(chan error)

		go processor.ProcessCollectionTask()(ctx, retryChan, []time.Time{})
		select {
		case err := <-retryChan:
			processor.logger.Infof("Processing error: \"%v\"", err)
		case <-ctx.Done():
			return
		}
	}

	notifications := processor.config.GetNotificationReceiversConfig().Notification
	if len(notifications) < 1 {
		processor.logger.Fatalf("Shutting down !!! No receivers in configuration !")
	}
	for _, notification := range notifications {
		err := processor.notifier.PublishToUser(ctx, processingError.Error(), notification.Receivers, ports.NotificationType(notification.NotificationType))
		if err != nil {
			return
		}
	}
}
