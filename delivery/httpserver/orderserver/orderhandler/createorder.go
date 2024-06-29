package orderhandler

import (
	"context"
	"fmt"
	"github.com/iam-benyamin/hellofresh/param/orderparam"
	"github.com/iam-benyamin/hellofresh/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) createOrder(c echo.Context) error {

	// TODO: validate request body
	var reqBody orderparam.CreateOrderRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Println("createOrderRequest")
	fmt.Println(reqBody)

	// TODO: the reqBody.ProductCode is actually a ProductID :D
	err = h.OrderService.CreateNewOrder(context.Background(), orderparam.CreateOrderRequest{
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
