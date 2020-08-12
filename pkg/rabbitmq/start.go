package rabbitmq

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	MqWait         *sync.WaitGroup
	ConsumeConn    *amqp.Connection
	ConsumeChannel *amqp.Channel
	PublishConn    *amqp.Connection
	PublishChannel *amqp.Channel
)

func Setup() {
	go ConsumeStart()
	go PublishStart()
}

func ConsumeStart() {
	ConsumeConnect()

	exchange := viper.GetString("ampq.consume.exchange")
	err := ConsumeChannel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println("rabbitmq ConsumeStart exchange: %v err: %v", exchange, err)
		ConsumeChannel.Close()
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}

	queue := viper.GetString("ampq.consume.queue")
	_, err = ConsumeChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		fmt.Println("rabbitmq ConsumeStart queue: %v err: %v", queue, err)
		ConsumeChannel.Close()
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}

	routingKey := viper.GetString("ampq.consume.queue")
	err = ConsumeChannel.QueueBind(queue, routingKey, exchange, false, nil)
	if err != nil {
		fmt.Println("rabbitmq ConsumeStart routingKey: %v err: %v", routingKey, err)
		ConsumeChannel.Close()
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		go ConsumeStart()
		return
	}

	fmt.Println("rabbitmq ConsumeStart 完成")

	MqWait = new(sync.WaitGroup)
	MqWait.Add(1)
	go MqConsume()
}

func ConsumeConnect() {
	fmt.Println("rabbitmq ConsumeConnect 开始连接")

	var err error

	rabbitmqUrl := fmt.Sprintf("amqp://%s:%s@%s/",
		viper.GetString("ampq.consume.username"),
		viper.GetString("ampq.consume.password"),
		viper.GetString("ampq.consume.addr"),
	)

do:
	ConsumeConn, err = amqp.Dial(rabbitmqUrl)
	if err != nil {
		fmt.Println("rabbitmq ConsumeConnect 连接失败 err: %v", err)
		time.Sleep(3 * time.Second)
		goto do
	}
	ConsumeChannel, err = ConsumeConn.Channel()
	if err != nil {
		fmt.Println("rabbitmq ConsumeConnect 打开channel失败 err = %v", err)
		ConsumeConn.Close()
		time.Sleep(3 * time.Second)
		goto do
	}

	fmt.Println("rabbitmq ConsumeConnect 连接完成")
}

func PublishStart() {
	PublishConnect()

	exchange := viper.GetString("ampq.publish.exchange")
	err := PublishChannel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println("rabbitmq PublishStart exchange: %v err: %v", exchange, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}

	queue := viper.GetString("ampq.publish.queue")
	_, err = PublishChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		fmt.Println("rabbitmq PublishStart queue: %v err: %v", queue, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}

	routingKey := viper.GetString("ampq.publish.queue")
	err = PublishChannel.QueueBind(queue, routingKey, exchange, false, nil)
	if err != nil {
		fmt.Println("rabbitmq ConsumeStart routingKey: %v err: %v", routingKey, err)
		PublishChannel.Close()
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		go PublishStart()
		return
	}

	fmt.Println("rabbitmq PublishStart 完成")
}

func PublishConnect() {
	fmt.Println("rabbitmq PublishConnect 开始连接")

	var err error

	rabbitmqUrl := fmt.Sprintf("amqp://%s:%s@%s/",
		viper.GetString("ampq.publish.username"),
		viper.GetString("ampq.publish.password"),
		viper.GetString("ampq.publish.addr"),
	)

do:
	PublishConn, err = amqp.Dial(rabbitmqUrl)
	if err != nil {
		fmt.Println("rabbitmq PublishConnect 连接失败 err: %v", err)
		time.Sleep(3 * time.Second)
		goto do
	}
	PublishChannel, err = PublishConn.Channel()
	if err != nil {
		fmt.Println("rabbitmq PublishConnect 打开channel失败 err = %v", err)
		PublishConn.Close()
		time.Sleep(3 * time.Second)
		goto do
	}

	fmt.Println("rabbitmq PublishConnect 连接完成")
}
