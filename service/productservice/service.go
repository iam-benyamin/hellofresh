package productservice

import (
	"context"

	"github.com/iam-benyamin/hellofresh/entity/productentity"
)

type Repository interface {
	GetProductByProductCode(ctx context.Context, ProductCode string) (productentity.Product, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
