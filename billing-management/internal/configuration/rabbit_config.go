/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
