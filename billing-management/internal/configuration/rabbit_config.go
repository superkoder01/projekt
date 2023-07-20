package configuration

import "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"

func GetRabbitMQConfig() *config.RabbitMQProducerConfig {

	ExchangeName := CS.GetString(string(ExchangeName))
	ExchangeType := CS.GetString(string(ExchangeType))
	QueueBindingsModel := CS.Get(string(QueueBindings))
	RabbitUrl := CS.GetString(string(RabbitUrl))
	IsDurable := CS.GetBool(string(IsDurable))
	IsAutoDelete := CS.GetBool(string(IsAutoDelete))
	var queueBindings []config.QueueBinding
	for _, b := range QueueBindingsModel.([]interface{}) {
		qbm := b.(map[interface{}]interface{})
		qb := &config.QueueBinding{
			BindingKey: qbm["BindingKey"].(string),
			QueueName:  []string{},
		}
		for _, qn := range qbm["QueueName"].([]interface{}) {
			qb.QueueName = append(qb.QueueName, qn.(string))
		}
		queueBindings = append(queueBindings, *qb)
	}
	mqConfig := config.RabbitMQConfig{
		RabbitUrl:    RabbitUrl,
		IsDurable:    IsDurable,
		IsAutoDelete: IsAutoDelete,
	}
	return &config.RabbitMQProducerConfig{
		ExchangeName:   ExchangeName,
		ExchangeType:   ExchangeType,
		QueueBindings:  queueBindings,
		RabbitMQConfig: mqConfig,
	}
}
