package server

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/email/consumer"
	"NotificationEmailService/internal/email/smtp"
	"NotificationEmailService/internal/email/usecase"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
	"go.uber.org/zap"
	"log"
	"strconv"
)

type Server struct {
	emailConfig *config.EmailServiceConfig
	logger      *zap.SugaredLogger
	smtp        smtp.Smtp
}

func New(config *config.EmailServiceConfig, logger *zap.SugaredLogger, smtp smtp.Smtp) *Server {
	return &Server{emailConfig: config, logger: logger, smtp: smtp}
}

func (server *Server) RunEmailServer() error {

	server.logger.Infof("%v is running ", server.emailConfig.Service.ServiceName)
	server.logger.Infof("Loaded configuration: %v", server.emailConfig.Service.ServiceName)

	consumerName := "emailConsumer"
	chanErr := make(chan error)

	emailUseCase := usecase.NewEmailUseCase(server.smtp, server.logger, conversion.NewByConfig(&server.emailConfig.Conversion, server.logger), &server.emailConfig.Service)
	emailConsumer := consumer.NewEmailConsumer(emailUseCase, consumerName)
	rabbit := rabbitmq.NewRabbitMQConsumer(server.emailConfig.Rabbitmq, emailConsumer, server.logger)

	err := rabbit.InitializeConnection()
	if err != nil {
		server.logger.Errorf("Failed to initialize rabbitMQ connection ! %v", err)
		return err
	}

	server.logger.Infof("EmailService components initialized")
	server.logger.Infof("Numbers of workers to create: " + strconv.Itoa(server.emailConfig.Service.WorkerPoolSize))

	for i := 1; i <= server.emailConfig.Service.WorkerPoolSize; i++ {
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
