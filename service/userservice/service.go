package userservice

import (
	"context"

	"github.com/iam-benyamin/hellofresh/entity"
)

type Repository interface {
	GetUserByID(ctx context.Context, UserID string) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
