package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/iam-benyamin/hellofresh/delivery/grpcserver/userserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqluser"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqluser/migrator"
	"github.com/iam-benyamin/hellofresh/service/userservice"
)

func main() {
	done := make(chan bool)
	wg := sync.WaitGroup{}

	// TODO: read all configs from file
	// TODO: logger
	cfg := mysql.Config{
		Host:     "localhost",
		Port:     3308,
		Username: "hellofresh",
		Password: "hellofresh0lk2o20",
		DBName:   "hellofresh_db",
	}

	mgr := migrator.New(cfg)
	mgr.Up()

	mysqlRepo := mysql.New(cfg)
	userMysql := mysqluser.New(mysqlRepo)

	userSVC := userservice.New(userMysql)

	server := userserver.New(userSVC)

	server.Start(done, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("\nReceived interrupt signal, shutting down gracefully...")

	done <- true
	close(done)

	wg.Wait()
	fmt.Println("Bay Bay")
}
