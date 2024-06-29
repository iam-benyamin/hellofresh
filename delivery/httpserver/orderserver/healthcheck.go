package orderserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (o OrderServer) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good",
	})
}
