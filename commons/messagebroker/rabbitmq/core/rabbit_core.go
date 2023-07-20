package core

import (
	"github.com/streadway/amqp"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/logger"
)

type rabbitMQCore struct {
	logger logger.Logger
	rabbitConnection
}

type rabbitConnection struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

/*	Close all channels and connections with RabbitMQ
	It is indicated to use this method on application shutdown
*/
func (rabbit *rabbitMQCore) Close() error {
	rabbit.logger.Infof("Shutting down rabbit connection")
	chanErr := rabbit.channel.Close()
	if chanErr != nil {
		return chanErr
	}
	connErr := rabbit.connection.Close()
	if connErr != nil {
		return connErr
	}

	return nil
}

func (rabbit *rabbitMQCore) createConnectionWithRabbitMQ(rabbitUrl string) error {
	connection, err := amqp.Dial(rabbitUrl)
	if err != nil {
		rabbit.logger.Errorf("Failed to connect to rabbitMQCore: %v", rabbitUrl)
		return err
	}

	channel, err := connection.Channel()
	if err != nil {
		rabbit.logger.Errorf("Failed to open a channel: %v", err)
		return err
	}
	rabbit.connection = connection
	rabbit.channel = channel

	return nil
}

func (rabbit *rabbitMQCore) registerChannel(queName, queType string, prefetchCount int, isDurable, isAutoDelete bool) (amqp.Queue, error) {
	queue, err := rabbit.channel.QueueDeclare(
		queName,
		isDurable,
		isAutoDelete,
		false,
		false,
		map[string]interface{}{"x-queue-type": queType},
	)
	if err != nil {
		rabbit.logger.Errorf("Failed to create queue: %v, of type: %v "+
			"with parameters: "+"isDurable:%v, isAutoDelete:%v, isExclusive:false, isNoWait:false",
			queName, queType, isDurable, isAutoDelete)
		return amqp.Queue{}, err
	}

	err = rabbit.channel.Qos(prefetchCount, 0, false)
	if err != nil {
		rabbit.logger.Errorf("Failed to set Qos with prefetchCount: %v", prefetchCount)
		return amqp.Queue{}, err
	}

	rabbit.logger.Infof("Created queue: %v, of type: %v "+
		"with parameters: "+"isDurable:%v, isAutoDelete:%v, isExclusive:false, isNoWait:false",
		queName, queType, isDurable, isAutoDelete)
	return queue, nil
}

func (rabbit *rabbitMQCore) createBindings(bindings []config.QueueBinding, exchangeName string) error {
	if len(bindings) >= 1 {
		var failCounter int
		var err error
		for _, binding := range bindings {
			for _, queue := range binding.QueueName {
				err = rabbit.bindQueueToExchange(queue, binding.BindingKey, exchangeName)
				if err != nil {
					failCounter++
				}
			}
		}
		if len(bindings) == failCounter {
			rabbit.logger.Errorf("Failed to bind any queue ! %v", err)
			return err
		} else if failCounter > 1 {
			rabbit.logger.Errorf("Bindings created but more than %v failed ! %v", failCounter, err)
			return err
		} else {
			rabbit.logger.Infof("Bindings created")
			return nil
		}
	}
	return nil
}

func (rabbit *rabbitMQCore) createExchange(exchangeName, exchangeType string, isDurable, isAutoDelete bool) error {
	err := rabbit.channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		isDurable,
		isAutoDelete,
		false,
		false,
		nil,
	)
	if err != nil {
		rabbit.logger.Errorf("Failed to create exchange: %v, of type: %v "+
			"with parameters: "+"isDurable:%v, isAutoDelete:%v, isExclusive:false, isNoWait:false",
			exchangeName, exchangeType, isDurable, isAutoDelete)
		return err
	}

	rabbit.logger.Infof("Created exchange: %v, of type: %v "+
		"with parameters: "+"isDurable:%v, isAutoDelete:%v, isExclusive:false, isNoWait:false",
		exchangeName, exchangeType, isDurable, isAutoDelete)
	return nil
}

func (rabbit *rabbitMQCore) bindQueueToExchange(queueName, routingKey, exchangeName string) error {
	err := rabbit.channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		rabbit.logger.Errorf("Failed to bind queue: %v, to exchange: %v", queueName, exchangeName)
		return err
	}

	rabbit.logger.Infof("Created queue binding: %v, to exchange: %v", queueName, exchangeName)
	return nil
}
