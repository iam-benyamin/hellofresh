package orderhandler

import "github.com/iam-benyamin/hellofresh/service/orderservice"

type Handler struct {
	OrderService orderservice.Service
	//	TODO: validator
}

func New(orderService orderservice.Service) Handler {
	return Handler{OrderService: orderService}
}
