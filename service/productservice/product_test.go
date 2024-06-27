package productservice_test

import (
	"context"
	"github.com/iam-benyamin/hellofresh/entity/productentity"
	"github.com/iam-benyamin/hellofresh/param/productparam"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/service/productservice"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

type InMemoryRepo struct {
	data map[string]productentity.Product
	mu   sync.Mutex
}

func InMemoryUserRepo() *InMemoryRepo {
	return &InMemoryRepo{
		data: make(map[string]productentity.Product),
	}
}

func (r *InMemoryRepo) GetProductByProductCode(ctx context.Context, userID string) (productentity.Product, error) {
	product, ok := r.data[userID]
	if !ok {
		return productentity.Product{}, richerror.New("productservice_test.GetProductByProductCode").WithMessage(errmsg.ErrorMsgNotFound).
			WithKind(richerror.KindNotFound)
	}
	return product, nil
}

func (r *InMemoryRepo) AddProduct(product productentity.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[product.Code] = product
}

func TestServiceProduct(t *testing.T) {
	repo := InMemoryUserRepo()
	service := productservice.New(repo)

	ctx := context.Background()

	repo.AddProduct(productentity.Product{ID: "1234", Name: "food1", Price: 25, Code: "family", CreatedAt: time.Now()})
	repo.AddProduct(productentity.Product{ID: "5678", Name: "food2", Price: 15, Code: "classic", CreatedAt: time.Now()})

	scenarios := []struct {
		name     string
		req      productparam.ProductRequest
		expected productparam.ProductResponse
		err      error
	}{
		{
			name: "success case - product found family",
			req:  productparam.ProductRequest{ProductCode: "family"},
			expected: productparam.ProductResponse{
				ID:    "1234",
				Name:  "food1",
				Code:  "family",
				Price: 25,
			},
			err: nil,
		},
		{
			name:     "error case - user fount 2468",
			req:      productparam.ProductRequest{ProductCode: "food-italy"},
			expected: productparam.ProductResponse{},
			err: richerror.New("productservice_test.GetUserByID").WithMessage(errmsg.ErrorMsgNotFound).
				WithKind(richerror.KindNotFound),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result, err := service.Product(ctx, scenario.req)

			if scenario.err != nil {
				assert.Error(t, err)
				assert.Equal(t, scenario.expected, result)
				assert.Equal(t, scenario.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, scenario.expected, result)
			}
		})
	}
}
