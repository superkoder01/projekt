package config

type RabbitMQConfig struct {
	RabbitUrl    string
	IsDurable    bool
	IsAutoDelete bool
}

type RabbitMQConsumerConfig struct {
	QueueName     string
	QueueType     string
	IsAutoAck     bool
	PrefetchCount int
	RabbitMQConfig
}

type RabbitMQProducerConfig struct {
	ExchangeName  string
	ExchangeType  string
	QueueBindings []QueueBinding
	RabbitMQConfig
}

type QueueBinding struct {
	QueueName  []string
	BindingKey string
}
