package userservice

import (
	"context"

	"github.com/iam-benyamin/hellofresh/entity/userentity"
)

type Repository interface {
	GetUserByID(ctx context.Context, UserID string) (userentity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
