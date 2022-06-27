package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang_learn/go_wire/internal/config"
)

// var Provider = wire.NewSet(New) // 同理

func NewDB(cfg *config.Config) (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		return
	}
	cleanup = func() {
		db.Close()
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, cleanup, nil
}
