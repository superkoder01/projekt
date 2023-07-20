package core

import (
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/logger"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/logger/zap"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/messagebroker/rabbitmq/consumer"
)

type rabbitMQConsumer struct {
	rabbitMQCore
	messageConsumer consumer.Worker
	config          config.RabbitMQConsumerConfig
	queue           amqp.Queue
}

/*
NewRabbitMQConsumer creates receiver which consumes messages from RabbitMQ.
If connection with RabbitMQ cannot be established due to configuration / external errors function calls panic.
Consumer use internal console logger, to customize logs see NewRabbitMQConsumerWithLogger
*/
func NewRabbitMQConsumer(mqConfig config.RabbitMQConsumerConfig, executor consumer.Worker) *rabbitMQConsumer {
	appLogger := zap.NewZapLogger(
		&config.LoggerConfig{
			Development: false,
			Encoding:    "console",
			Level:       "debug",
		})

	consumer := rabbitMQConsumer{
		messageConsumer: executor,
		config:          mqConfig,
		rabbitMQCore: rabbitMQCore{
			logger: zap.CreateLoggerContext(*appLogger, "RabbitMQ-Consumer"),
		},
	}

	consumer.initializeConnection()
	return &consumer
}

/*
NewRabbitMQConsumerWithLogger creates receiver which consumes messages from RabbitMQ.
If connection with RabbitMQ cannot be established due to configuration / external errors function calls panic.
Consumer use injected logger.Logger implementation
*/
func NewRabbitMQConsumerWithLogger(mqConfig config.RabbitMQConsumerConfig, executor consumer.Worker, logger logger.Logger) *rabbitMQConsumer {
	consumer := rabbitMQConsumer{
		messageConsumer: executor,
		config:          mqConfig,
		rabbitMQCore: rabbitMQCore{
			logger: logger,
		},
	}

	consumer.initializeConnection()
	return &consumer
}

func (rabbit *rabbitMQConsumer) initializeConnection() {
	if len(rabbit.config.QueueName) < 1 || len(rabbit.config.QueueType) < 1 {
		rabbit.logger.Fatalf("Invalid configuration !! Queue must be declared")
		return
	}

	connErr := rabbit.rabbitMQCore.createConnectionWithRabbitMQ(rabbit.config.RabbitUrl)
	if connErr != nil {
		rabbit.logger.Fatalf("Failed to create connection with rabbitMq ! %v", connErr)
		return
	} else {
		rabbit.logger.Infof("Successfully connected to rabbitMQCore service")
	}

	queue, chanelErr := rabbit.registerChannel(
		rabbit.config.QueueName,
		rabbit.config.QueueType,
		rabbit.config.PrefetchCount,
		rabbit.config.IsDurable,
		rabbit.config.IsAutoDelete,
	)
	if chanelErr != nil {
		rabbit.logger.Fatalf("Failed to create channel ! %v", chanelErr)
		return
	}

	rabbit.queue = queue
	rabbit.logger.Debugf("RabbitConsumer initialized")
}

/* StartConsumer create consumer listening on rabbitMQ channel and create goroutine with consumer.Worker*/
func (rabbit *rabbitMQConsumer) StartConsumer(consumerName string) error {
	payload, err := rabbit.registerConsumer(consumerName)
	if err != nil {
		return err
	}

	rabbit.logger.Infof("[%v] Listening for events...", consumerName)
	go rabbit.worker(payload, consumerName)

	chErr := <-rabbit.channel.NotifyClose(make(chan *amqp.Error))
	rabbit.logger.Errorf("ch.NotifyClose: %v", chErr)
	return chErr
}

func (rabbit *rabbitMQConsumer) registerConsumer(consumerName string) (<-chan amqp.Delivery, error) {
	payload, err := rabbit.channel.Consume(
		rabbit.queue.Name,
		consumerName,
		rabbit.config.IsAutoAck,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		rabbit.logger.Errorf("%v: failed to register rabbit rabbitMQService %v", consumerName, err)
		return nil, err
	}

	rabbit.logger.Infof("Consumer %v created", consumerName)
	return payload, nil
}

func (rabbit *rabbitMQConsumer) worker(payload <-chan amqp.Delivery, consumerName string) {
	apiLogger := rabbit.logger
	apiLogger.Infof("[%v] Worker started", consumerName)
	for message := range payload {
		apiLogger.Infof("[%v] Worker is processing message: %v", consumerName, string(message.Body))

		messageId := message.MessageId
		if len(messageId) < 1 {
			messageId = uuid.New().String()
			apiLogger.Infof("[%v] Received message with no identifier !! Generated new id %v for message: %v", consumerName, messageId, string(message.Body[:]))
		}

		err := rabbit.messageConsumer.Execute(message.Body, messageId, consumerName)

		if err == nil {
			apiLogger.Infof("[%v] Message[%v] processing done !", consumerName, messageId)
			rabbit.sendAckMessage(message, consumerName+":"+messageId)
		} else {
			rabbit.handleMessageError(message, err, consumerName+":"+messageId)
		}
	}
	apiLogger.Errorf("[%v] Delivery channel closed", consumerName)
}

/*todo add requeueing to different que or limited for 3*/
func (rabbit *rabbitMQConsumer) handleMessageError(message amqp.Delivery, error error, name string) {
	/*if errors.As(error, &customError.NotReadyMessage{}) {
		err := message.Nack(false, true)
		if err != nil {
			rabbit.logger.Errorf("%v: failed to acknowledge deliver %v %v", name, message, err)
		}
		rabbit.logger.Errorf("Service unable to process message - re-queuing !!")
	} else if errors.As(error, &customError.IncorrectMessage{}) {
		rabbit.logger.Errorf("Message cannot be processed ! Skipping !!")
		rabbit.sendAckMessage(message, name)
	} else if errors.As(error, &customError.MessageServiceError{}) {
		rabbit.logger.Fatalf("Fatal error encountered ! Shutting down !!")
	} else {
		rabbit.logger.Errorf("%v: Unknown error - cannot execute task !! ", name)
		rabbit.sendAckMessage(message, name)
	}*/
	rabbit.logger.Errorf("Error: %v", error)
}

func (rabbit *rabbitMQConsumer) sendAckMessage(message amqp.Delivery, name string) {
	ackErr := message.Ack(false)
	if ackErr != nil {
		rabbit.logger.Errorf("%v: failed to acknowledge deliver %v", name, ackErr)
	}
}
