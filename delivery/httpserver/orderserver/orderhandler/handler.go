package orderhandler

import (
	"github.com/iam-benyamin/hellofresh/service/orderservice"
	"github.com/iam-benyamin/hellofresh/validator/ordervaidator"
)

type Handler struct {
	OrderService         orderservice.Service
	CreateOrderValidator ordervaidator.Validator
}

func New(orderService orderservice.Service, createOrderValidator ordervaidator.Validator) Handler {
	return Handler{OrderService: orderService, CreateOrderValidator: createOrderValidator}
}
