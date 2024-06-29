package orderserver

import (
	"fmt"
	"github.com/iam-benyamin/hellofresh/delivery/httpserver/orderserver/orderhandler"
	"github.com/iam-benyamin/hellofresh/logger"
	"github.com/iam-benyamin/hellofresh/service/orderservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type OrderServer struct {
	Router       *echo.Echo
	OrderHandler orderhandler.Handler
}

func New(OrderService orderservice.Service) OrderServer {
	return OrderServer{
		Router:       echo.New(),
		OrderHandler: orderhandler.New(OrderService),
	}
}

func (o OrderServer) Serve() {
	o.Router.Use(middleware.RequestID())
	o.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			logger.Logger.Named("order-http-server").Info("request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	o.Router.Use(middleware.Recover())

	o.Router.GET("/heath-check/", o.healthCheck)
	o.OrderHandler.SetOrderRoutes(o.Router)

	address := fmt.Sprintf(":%d", 1323)
	if err := o.Router.Start(address); err != nil {
		fmt.Println("router start error ", err)
	}
}
