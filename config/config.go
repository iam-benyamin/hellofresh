package config

import (
	"github.com/iam-benyamin/hellofresh/adapter/rabbitmqadapter"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
)

type Config struct {
	UserDB    mysql.Config           `koanf:"user_db"`
	OrderDB   mysql.Config           `koanf:"order_db"`
	ProductDB mysql.Config           `koanf:"product_db"`
	OrderRMQ  rabbitmqadapter.Config `koanf:"order_rmq"`
}
