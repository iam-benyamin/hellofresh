package main

import (
	"fmt"
	"github.com/iam-benyamin/hellofresh/delivery/httpserver/orderserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqlorder"
	"github.com/iam-benyamin/hellofresh/service/orderservice"
	"os"
	"os/signal"
	"sync"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	cfg := mysql.Config{
		Host:     "localhost",
		Port:     3309,
		Username: "hellofresh",
		Password: "orderPassword",
		DBName:   "order_db",
	}

	dialect := "mysql"
	migrationPath := "repository/mysql/mysqlorder/migrations"

	mgr := migrator.New(dialect, cfg, migrationPath)
	//mgr.Down()
	mgr.Up()

	mysqlRepo := mysql.New(cfg)
	orderMysql := mysqlorder.New(mysqlRepo)

	orderSVC := orderservice.New(orderMysql)

	server := orderserver.New(orderSVC)
	server.Serve(done, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\nReceived interrupt signal, shutting user service down gracefully...")

	done <- true
	close(done)

	wg.Wait()
	fmt.Println("I hope to see you soon")
}
