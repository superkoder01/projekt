package server

import (
	"NotificationSmsService/internal/sms/consumer"
	"NotificationSmsService/internal/sms/gateway"
	"NotificationSmsService/internal/sms/usecase"
	"NotificationSmsService/pkg/config"
	"NotificationSmsService/pkg/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
	"log"
	"strconv"
)

type Server struct {
	smsConfig *config.SmsServiceConfig
	logger    logger.Logger
}

func NewSmsServer(config *config.SmsServiceConfig, logger logger.Logger) *Server {
	return &Server{smsConfig: config, logger: logger}
}

func (server *Server) RunSmsServer() error {
	consumerName := "smsConsumer"

	chanErr := make(chan error)
	smsGateway := gateway.NewGateway(server.smsConfig.Gateway, server.logger)
	smsUseCase := usecase.NewSmsUseCase(*smsGateway, server.logger)
	smsConsumer := consumer.NewSmsConsumer(smsUseCase, consumerName)
	rabbit := rabbitmq.NewRabbitMQConsumer(server.smsConfig.Rabbitmq, smsConsumer, server.logger)

	err := rabbit.InitializeConnection()
	if err != nil {
		server.logger.Errorf("Failed to initialize rabbitMQ connection ! %v", err)
		return err
	}

	server.logger.Infof("SmsService components initialized")
	server.logger.Infof("Numbers of workers to create: " + strconv.Itoa(server.smsConfig.Service.WorkerPoolSize))

	for i := 1; i <= server.smsConfig.Service.WorkerPoolSize; i++ {
		go func(index string) {
			err := rabbit.StartConsumer(consumerName + "-" + index)
			if err != nil {
				chanErr <- err
			}
		}(strconv.Itoa(i))
	}

	log.Printf("Error %v", <-chanErr)
	_ = rabbit.Close()
	return nil
}
