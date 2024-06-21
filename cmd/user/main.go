package main

import (
	"github.com/iam-benyamin/hellofresh/delivery/grpcserver/userserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqluser"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqluser/migrator"
	"github.com/iam-benyamin/hellofresh/service/userservice"
)

func main() {
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
	server.Start()
}
