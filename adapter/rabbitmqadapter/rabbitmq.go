package rabbitmqadapter

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Exchange string
}

type Config struct {
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
}

func New(config Config, exchange string) (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.User, config.Password, config.Host, config.Port))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()

		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		_ = ch.Close()
		_ = conn.Close()

		return nil, err
	}

	return &RabbitMQAdapter{
		Conn:     conn,
		Channel:  ch,
		Exchange: exchange,
	}, nil
}

func (r *RabbitMQAdapter) Close() {
	if r.Channel != nil {
		_ = r.Channel.Close()
	}
	if r.Conn != nil {
		_ = r.Conn.Close()
	}
}
