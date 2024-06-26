package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/iam-benyamin/hellofresh/config"
	"github.com/iam-benyamin/hellofresh/delivery/grpcserver/productserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqlproduct"
	"github.com/iam-benyamin/hellofresh/service/productservice"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	cfg := config.Load("config.yaml")

	// TODO: read all configs from file
	// TODO: logger
	dialect := "mysql"
	migrationPath := "repository/mysql/mysqlproduct/migrations"

	mgr := migrator.New(dialect, cfg.ProductDB, migrationPath)
	// mgr.Down()
	mgr.Up()

	mysqlRepo := mysql.New(cfg.ProductDB)
	productMysql := mysqlproduct.New(mysqlRepo)

	productSVC := productservice.New(productMysql)

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
