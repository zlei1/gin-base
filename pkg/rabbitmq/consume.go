package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/spf13/viper"

	"gin-base/app/workers"
)

func MqConsume() {
	if ConsumeChannel == nil {
		ConsumeConnect()
	}

	err := ConsumeChannel.Qos(1, 0, true)
	if err != nil {
		log.Fatalf("MqConsume Qos Failed: %s", err)
	}

	queue := viper.GetString("ampq.consume.queue")
	msgs, err := ConsumeChannel.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("MqConsume Failed to register a consumer: %s", err)

		return
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Printf("MqConsume Received a message: %s", msg.Body)

			payload := make(map[string]interface{})
			json.Unmarshal(msg.Body, &payload)

			workers.Assign(payload)

			msg.Ack(false)
		}
	}()

	log.Printf("MqConsume Waiting for messages...")
	<-forever
}
