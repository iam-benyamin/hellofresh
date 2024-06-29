package main

import (
	"github.com/iam-benyamin/hellofresh/delivery/httpserver/orderserver"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	"github.com/iam-benyamin/hellofresh/repository/mysql/migrator"
	"github.com/iam-benyamin/hellofresh/repository/mysql/mysqlorder"
	"github.com/iam-benyamin/hellofresh/service/orderservice"
)

func main() {
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

	// start order server
	server := orderserver.New(orderSVC)
	server.Serve()

	// TODO: don't forget remove this faker
	//fakeRequest := orderparam.CreateOrderRequest{
	//	UserID:      "b7c8d9e0f1a2",
	//	ProductCode: "b2c3d4e5f6g7",
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//
	//err := orderSVC.CreateNewOrder(ctx, fakeRequest)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// get user profile with UserID

}