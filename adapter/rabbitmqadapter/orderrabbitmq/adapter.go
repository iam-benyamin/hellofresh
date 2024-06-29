package orderrabbitmq

import "github.com/iam-benyamin/hellofresh/adapter/rabbitmqadapter"

type Adapter struct {
	rmq *rabbitmqadapter.RabbitMQAdapter
}

func New(rmq *rabbitmqadapter.RabbitMQAdapter) *Adapter {
	return &Adapter{rmq: rmq}
}
