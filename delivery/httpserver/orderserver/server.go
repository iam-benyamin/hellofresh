package orderserver

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/iam-benyamin/hellofresh/delivery/httpserver/orderserver/orderhandler"
	"github.com/iam-benyamin/hellofresh/logger"
	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/service/orderservice"
	"github.com/iam-benyamin/hellofresh/validator/ordervaidator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type OrderServer struct {
	Router       *echo.Echo
	OrderHandler orderhandler.Handler
}

func New(orderService orderservice.Service, createOrderValidator ordervaidator.Validator) OrderServer {
	return OrderServer{
		Router:       echo.New(),
		OrderHandler: orderhandler.New(orderService, createOrderValidator),
	}
}

func (o OrderServer) Serve(done <-chan bool, wg *sync.WaitGroup) {
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
	o.Router.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: errmsg.ErrorMsgSomeThingWentWrong,
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			logger.Logger.Error("time Out")
		},
		Timeout: 3 * time.Second,
	}))

	o.Router.GET("/heath-check/", o.healthCheck)
	o.OrderHandler.SetOrderRoutes(o.Router)

	go func() {
		address := fmt.Sprintf(":%d", 1323)
		if err := o.Router.Start(address); err != nil {
			fmt.Println("router start error ", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-done
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := o.Router.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
		fmt.Println("order http server shutdown gracefully")
	}()
}
