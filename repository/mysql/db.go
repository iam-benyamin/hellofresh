package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int16  `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	DBName   string `koanf:"db_name"`
}

type MySQLDB struct {
	db     *sql.DB
	config Config
}

func New(config Config) *MySQLDB {
	// TODO: implement retry policy to check if DB is ready or not
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName,
	))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db : %w", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{config: config, db: db}
}

func (m *MySQLDB) Conn() *sql.DB {
	return m.db
}
