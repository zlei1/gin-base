package rabbitmq

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
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
		log.Fatalf("%s: %s", "ConsumeStart Failed to declare exchange", err)
	}

	queue := viper.GetString("ampq.consume.queue")
	_, err = ConsumeChannel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "ConsumeStart Failed to declare queue", err)
	}

	routingKey := viper.GetString("ampq.consume.queue")
	err = ConsumeChannel.QueueBind(queue, routingKey, exchange, false, nil)
	if err != nil {
		log.Fatalf("%s: %s", "ConsumeStart Failed to bind queue", err)
	}

	fmt.Println("rabbitmq ConsumeStart 完成")

	go MqConsume()
}

func ConsumeConnect() {
	log.Println("ConsumeConnect 开始连接")

	var err error

	rabbitmqUrl := fmt.Sprintf("amqp://%s:%s@%s/",
		viper.GetString("ampq.consume.username"),
		viper.GetString("ampq.consume.password"),
		viper.GetString("ampq.consume.addr"),
	)

	ConsumeConn, err = amqp.Dial(rabbitmqUrl)
	if err != nil {
		log.Fatalf("%s: %s", "ConsumeConnect Failed to connect to RabbitMQ", err)
	}

	ConsumeChannel, err = ConsumeConn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "ConsumeConnect Failed to open a channel", err)
	}

	log.Println("ConsumeConnect 连接完成")
}

func PublishStart() {
	PublishConnect()

	exchange := viper.GetString("ampq.publish.exchange")
	err := PublishChannel.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "PublishStart Failed to declare exchange", err)
	}

	queue := viper.GetString("ampq.publish.queue")
	_, err = PublishChannel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "PublishStart Failed to declare queue", err)
	}

	routingKey := viper.GetString("ampq.publish.routingKey")
	err = PublishChannel.QueueBind(
		queue,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", "PublishStart Failed to bind queue", err)
	}

	log.Println("rabbitmq PublishStart 完成")
}

func PublishConnect() {
	log.Println("PublishConnect 开始连接")

	var err error

	rabbitmqUrl := fmt.Sprintf("amqp://%s:%s@%s/",
		viper.GetString("ampq.publish.username"),
		viper.GetString("ampq.publish.password"),
		viper.GetString("ampq.publish.addr"),
	)

	PublishConn, err = amqp.Dial(rabbitmqUrl)
	if err != nil {
		log.Fatalf("%s: %s", "PublishConnect Failed to connect to RabbitMQ", err)
	}

	PublishChannel, err = PublishConn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "PublishConnect Failed to open a channel", err)
	}

	log.Println("PublishConnect 连接完成")
}
