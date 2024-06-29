package ordervaidator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/iam-benyamin/hellofresh/param/orderparam"
)

func (v Validator) ValidateCrateOrderRequest(req orderparam.CreateOrderRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.UserID, validation.Required, validation.Length(12, 12)),
		validation.Field(&req.ProductCode, validation.Required),
	)
}
