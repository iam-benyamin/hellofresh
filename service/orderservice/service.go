package orderservice

import (
	"context"
	"github.com/iam-benyamin/hellofresh/param/orderparam"
)

type repository interface {
	SaveOrder(ctx context.Context, createOrder orderparam.SaveOrder) error
}

type Service struct {
	repo repository
}

func New(repo repository) Service {
	return Service{repo: repo}
}
