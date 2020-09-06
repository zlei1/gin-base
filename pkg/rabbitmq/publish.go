package rabbitmq

import (
	"encoding/json"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"gin-base/pkg/log"
)

type Message struct {
	Payload interface{} `json:"payload"`
}

func MqPublish(body *Message) {
	if PublishChannel == nil {
		PublishStart()
	}

	exchange := viper.GetString("ampq.publish.exchange")
	routingKey := viper.GetString("ampq.publish.routingKey")

	payload, _ := json.Marshal(body.Payload)
	err := PublishChannel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        payload,
	})

	if err != nil {
		log.Error("MqPublish Failed: %s", err)

		return
	}

	log.Info("MqPublish Succeed")
}
