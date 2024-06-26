package main

import (
	"fmt"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"os"
	"os/signal"
	"sync"

	"github.com/iam-benyamin/hellofresh/delivery/grpcserver/userserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqluser"
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
		Password: "userPassword",
		DBName:   "user_db",
	}

	dialect := "mysql"
	migrationPath := "repository/mysql/mysqluser/migrations"

	mgr := migrator.New(dialect, cfg, migrationPath)
	//mgr.Down()
	mgr.Up()

	mysqlRepo := mysql.New(cfg)
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
