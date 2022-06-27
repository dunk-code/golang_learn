package model

import "database/sql"

type App struct { // 最终需要的对象
	Db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{Db: db}
}
