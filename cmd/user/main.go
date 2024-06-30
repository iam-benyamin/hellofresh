package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/iam-benyamin/hellofresh/config"
	"github.com/iam-benyamin/hellofresh/delivery/grpcserver/userserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqluser"
	"github.com/iam-benyamin/hellofresh/service/userservice"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	// TODO: read all configs from file
	// TODO: logger
	cfg := config.Load("config.yaml")

	dialect := "mysql"
	migrationPath := "repository/mysql/mysqluser/migrations"

	mgr := migrator.New(dialect, cfg.UserDB, migrationPath)
	// mgr.Down()
	mgr.Up()

	mysqlRepo := mysql.New(cfg.UserDB)
	userMysql := mysqluser.New(mysqlRepo)

	userSVC := userservice.New(userMysql)

	server := userserver.New(userSVC)

	server.Start(done, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\nReceived interrupt signal, shutting user service down gracefully...")

	done <- true
	close(done)

	wg.Wait()
	fmt.Println("I hope to see you soon")
}
