package userservice

import (
	"context"
	"github.com/iam-benyamin/hellofresh/param/userparam"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
)

func (s Service) Profile(ctx context.Context, req userparam.ProfileRequest) (userparam.ProfileResponse, error) {
	const op = "userservice.Profile"

	u, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {

		return userparam.ProfileResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindNotFound).
			WithMeta(map[string]interface{}{"req": req})
	}

	return userparam.ProfileResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}, nil
}
