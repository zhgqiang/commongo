package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// RabbitMQ is 配置信息.
type RabbitMQ struct {
	Host       string `json:"host" toml:"host" description:"消息队列地址"`
	Port       int    `json:"port" toml:"port" description:"消息队列端口"`
	Username   string `json:"username" toml:"username" description:"消息队列访问用户名"`
	Password   string `json:"password" toml:"password" description:"消息队列访问密码"`
	Kind       string `json:"kind" toml:"kind" description:"消息队列访问密码"`
	VHost      string `json:"vHost" toml:"vHost" description:"消息队列VHost名"`
	Exchange   string `json:"exchange" toml:"exchange" description:"消息队列Exchange名"`
	RoutingKey string `json:"routingKey" toml:"routingKey" description:"消息队列名"`
}

// Send is RabbitMQ 发送信息.
func (p *RabbitMQ) Send(msg string) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s", p.Username, p.Password, p.Host, p.Port, p.VHost))
	if err != nil {
		return fmt.Errorf("RabbitMQ连接失败.%v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("RabbitMQ管道打开失败.%v", err)
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		p.Exchange, // name
		p.Kind,     // type
		false,      // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return fmt.Errorf("RabbitMQ的exchange定义失败.%v", err)
	}
	err = ch.Publish(
		p.Exchange,   // exchange
		p.RoutingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Expiration:  "60000",
			Body:        []byte(msg),
		})
	if err != nil {
		return fmt.Errorf("RabbitMQ发送消息失败.%v", err)
	}
	return nil
}
