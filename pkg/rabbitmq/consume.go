package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"gin-base/pkg/log"
)

var ConsumeWait *sync.WaitGroup

func MqConsume() {
	if ConsumeChannel == nil {
		ConsumeConnect()
	}

	err := ConsumeChannel.Qos(1, 0, true)
	if err != nil {
		log.Error("MqConsume Qos失败 err: %v", err)

		return
	}

	queue := viper.GetString("ampq.consume.queue")
	msgs, err := ConsumeChannel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		log.Error("MqConsume 接收消息失败 err: %v", err)

		return
	}

	ConsumeWait = new(sync.WaitGroup)
	for i := 0; i < viper.GetInt("ampq.consume.channel_range"); i++ {
		ConsumeWait.Add(1)
		go rangeChannel(msgs)
	}
	ConsumeWait.Wait()
}

func rangeChannel(msgs <-chan amqp.Delivery) {
	defer ConsumeWait.Done()
	for msg := range msgs {
		fmt.Println("msg: %v", msg.Body)

		msg.Ack(false)
	}
}
