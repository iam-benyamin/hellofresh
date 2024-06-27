package productservice

import (
	"context"

	"github.com/iam-benyamin/hellofresh/param/productparam"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
)

func (s Service) Product(ctx context.Context, req productparam.ProductRequest) (productparam.ProductResponse, error) {
	const op = "productservice.Product"

	p, err := s.repo.GetProductByProductCode(ctx, req.ProductCode)
	if err != nil {
		return productparam.ProductResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindNotFound).
			WithMeta(map[string]interface{}{"req": req})
	}

	return productparam.ProductResponse{
		ID:    p.ID,
		Name:  p.Name,
		Code:  p.Code,
		Price: p.Price,
	}, nil
}
