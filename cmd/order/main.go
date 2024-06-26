package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/iam-benyamin/hellofresh/adapter/rabbitmqadapter"
	"github.com/iam-benyamin/hellofresh/adapter/rabbitmqadapter/orderrabbitmq"
	"github.com/iam-benyamin/hellofresh/config"
	"github.com/iam-benyamin/hellofresh/delivery/httpserver/orderserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqlorder"
	"github.com/iam-benyamin/hellofresh/service/orderservice"
	"github.com/iam-benyamin/hellofresh/validator/ordervaidator"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	cfg := config.Load("config.yaml")

	dialect := "mysql"
	migrationPath := "repository/mysql/mysqlorder/migrations"

	mgr := migrator.New(dialect, cfg.OrderDB, migrationPath)
	// mgr.Down()
	mgr.Up()

	mysqlRepo := mysql.New(cfg.OrderDB)
	orderMysql := mysqlorder.New(mysqlRepo)

	rabbitmqAdapter, err := rabbitmqadapter.New(cfg.OrderRMQ, "orders")
	if err != nil {
		panic(err)
	}
	broker := orderrabbitmq.New(rabbitmqAdapter)

	orderSVC := orderservice.New(orderMysql, broker)

	validator := ordervaidator.New()

	server := orderserver.New(orderSVC, validator)
	server.Serve(done, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\nReceived interrupt signal, shutting user service down gracefully...")

	rabbitmqAdapter.Close()

	done <- true
	close(done)

	wg.Wait()
	fmt.Println("I hope to see you soon")
}
