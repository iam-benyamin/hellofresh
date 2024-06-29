package orderhandler

import "github.com/labstack/echo/v4"

func (h Handler) SetOrderRoutes(e *echo.Echo) {
	orderGroup := e.Group("/orders")

	orderGroup.POST("/", h.createOrder)
}
