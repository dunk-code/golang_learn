package main

import (
	"golang_learn/go_wire/wire"
	"log"
)

func main() {
	app, cleanup, err := wire.InitApp() // 使用wire生成的injector方法获取app对象
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	var version string
	row := app.Db.QueryRow("SELECT VERSION()")
	if err := row.Scan(&version); err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}
