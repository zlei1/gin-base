package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"gin-base/pkg/log"
)

func MqPublish(msg interface{}) {
	if PublishChannel == nil {
		PublishStart()
	}

	publishContent, err := json.Marshal(msg)
	if err != nil {
		log.Error("MqPublish Marshal err: %v", err)

		return
	}

	exchange := viper.GetString("ampq.publish.exchange")
	routingKey := viper.GetString("ampq.publish.routingKey")

	err = PublishChannel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        publishContent,
	})

	if err != nil {
		log.Error("MqPublish err: %v", err)

		return
	}

	fmt.Println("MqPublish success")
}
