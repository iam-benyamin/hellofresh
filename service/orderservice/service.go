package orderservice

import (
	"context"

	"github.com/iam-benyamin/hellofresh/param/orderparam"
)

type repository interface {
	SaveOrder(ctx context.Context, createOrder orderparam.SaveOrder) (string, error)
}

type Broker interface {
	PublishCreatedOrder(ctx context.Context, msg orderparam.Message, routingKey string) error
}

type Service struct {
	repo   repository
	broker Broker
}

func New(repo repository, broker Broker) Service {
	return Service{repo: repo, broker: broker}
}
