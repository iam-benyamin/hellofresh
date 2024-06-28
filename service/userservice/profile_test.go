package userservice_test

import (
	"context"
	"sync"
	"testing"

	"github.com/iam-benyamin/hellofresh/entity/userentity"
	"github.com/iam-benyamin/hellofresh/param/userparam"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/service/userservice"
	"github.com/stretchr/testify/assert"
)

type InMemoryRepo struct {
	data map[string]userentity.User
	mu   sync.Mutex
}

func InMemoryUserRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data: make(map[string]userentity.User),
	}
}

func (r *InMemoryRepo) GetUserByID(_ context.Context, userID string) (userentity.User, error) {
	user, ok := r.data[userID]
	if !ok {
		return userentity.User{}, richerror.New("userservice_test.GetUserByID").WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound)
	}
	return user, nil
}

func (r *InMemoryRepo) AddUser(user userentity.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[user.ID] = user
}

func TestServiceProfile(t *testing.T) {
	repo := InMemoryUserRepo()
	service := userservice.New(repo)

	ctx := context.Background()

	repo.AddUser(userentity.User{ID: "123", FirstName: "Milad", LastName: "Ahmadi"})
	repo.AddUser(userentity.User{ID: "456", FirstName: "Sara", LastName: "Moradi"})

	scenarios := []struct {
		name     string
		req      userparam.ProfileRequest
		expected userparam.ProfileResponse
		err      error
	}{
		{
			name: "success case - user 123",
			req:  userparam.ProfileRequest{UserID: "123"},
			expected: userparam.ProfileResponse{
				ID:        "123",
				FirstName: "Milad",
				LastName:  "Ahmadi",
			},
			err: nil,
		},
		{
			name:     "error case - user not found",
			req:      userparam.ProfileRequest{UserID: "999"},
			expected: userparam.ProfileResponse{},
			err: richerror.New("userservice.Profile").WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound).
				WithMeta(map[string]interface{}{"req": userparam.ProfileRequest{UserID: "999"}}),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			resp, err := service.Profile(ctx, scenario.req)

			if scenario.err != nil {
				assert.Error(t, err)
				assert.Equal(t, scenario.expected, resp)
				assert.Equal(t, scenario.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, scenario.expected, resp)
			}
		})
	}
}
