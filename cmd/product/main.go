package main

import (
	"fmt"
	"github.com/iam-benyamin/hellofresh/delivery/grpcserver/productserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqlproduct"
	"github.com/iam-benyamin/hellofresh/service/productservice"
	"os"
	"os/signal"
	"sync"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	// TODO: read all configs from file
	// TODO: logger
	cfg := mysql.Config{
		Host:     "localhost",
		Port:     3307,
		Username: "hellofresh",
		Password: "productPassword",
		DBName:   "product_db",
	}
	dialect := "mysql"
	migrationPath := "repository/mysql/mysqlproduct/migrations"

	mgr := migrator.New(dialect, cfg, migrationPath)
	mgr.Down()
	mgr.Up()

	mysqlRepo := mysql.New(cfg)
	productMysql := mysqlproduct.New(mysqlRepo)

	productSVC := productservice.New(productMysql)

	// TODO: server
	server := productserver.New(productSVC)

	server.Start(done, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\nReceived interrupt signal, shutting down product service gracefully...")

	done <- true
	close(done)

	wg.Wait()
	fmt.Println("Have go time ;)")
}
