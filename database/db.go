package database

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

var once sync.Once
var erro error
var db *sql.DB

func GetDb(ctx context.Context) (*sql.DB, error) {
	once.Do(func() {
		cfg := mysql.NewConfig()

		cfg.User = "root"
		cfg.Addr = "mysql-server"
		cfg.ParseTime = true
		cfg.Passwd = "123456"
		cfg.DBName = "todo"

		ctor, err := mysql.NewConnector(cfg)
		if err != nil {
			erro = err
			return
		}

		db = sql.OpenDB(ctor)

		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Minute * 3)

		err = db.PingContext(ctx)
		if err != nil {
			erro = err
		}
	})

	return db, erro
}
