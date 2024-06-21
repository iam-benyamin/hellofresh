package migrator

import (
	"database/sql"
	"fmt"
	"github.com/iam-benyamin/hellofresh/repository/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   mysql.Config
	migrations *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config) Migrator {
	return Migrator{
		dialect:    "mysql",
		dbConfig:   dbConfig,
		migrations: &migrate.FileMigrationSource{Dir: "repository/mysql/mysqluser/migrations"},
	}
}

func (m Migrator) Up() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName,
	))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't appliy migrations: %v", err))
	}

	fmt.Printf("Applied %d migrations\n", n)
}

func (m Migrator) Down() {
	db, err := sql.Open(m.dialect, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName,
	))
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't roll back migrations: %v", err))
	}

	fmt.Printf("RollBacked %d migrations!\n", n)
}
