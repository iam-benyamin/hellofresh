package orderhandler

import (
	"context"
	"net/http"
	"time"

	"github.com/iam-benyamin/hellofresh/param/orderparam"
	"github.com/iam-benyamin/hellofresh/pkg/httpmsg"
	"github.com/labstack/echo/v4"
)

func (h Handler) createOrder(c echo.Context) error {
	var reqBody orderparam.CreateOrderRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := h.CreateOrderValidator.ValidateCrateOrderRequest(reqBody); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO: the reqBody.ProductCode is actually a ProductID :D
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()
	err = h.OrderService.CreateNewOrder(ctx, orderparam.CreateOrderRequest{
		UserID:      reqBody.UserID,
		ProductCode: reqBody.ProductCode,
	})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "order created!",
	})
}
