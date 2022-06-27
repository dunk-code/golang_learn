package config

import (
	"encoding/json"
	"os"
)

// var Provider = wire.NewSet(New) // 将New方法声明为Provider，表示New方法可以创建一个被别人依赖的对象,也就是Config对象

type Config struct {
	Database database `json:"database"`
}

type database struct {
	Dsn string `json:"dsn"`
}

func NewConfig() (*Config, error) {
	fp, err := os.Open("go_wire/config/app.json")
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var cfg Config
	if err := json.NewDecoder(fp).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
