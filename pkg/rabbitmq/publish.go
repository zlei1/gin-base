package rabbitmq

import (
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"gin-base/pkg/log"
)

func MqPublish(body string) {
	if PublishChannel == nil {
		PublishStart()
	}

	exchange := viper.GetString("ampq.publish.exchange")
	routingKey := viper.GetString("ampq.publish.routingKey")

	err := PublishChannel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	if err != nil {
		log.Error("MqPublish Failed: %s", err)

		return
	}

	log.Info("MqPublish Succeed")
}
