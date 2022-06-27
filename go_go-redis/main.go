package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
	keys := make([]string, 2)
	keys[0] = "user_8648702"
	keys[1] = "user_6555122"
	vals, err := client.MGet(keys...).Result()
	if err != nil {
		panic(err)
	}
	for _, v := range vals {
		s, ok := v.(string)
		if ok {
			fmt.Println(s)
		}
	}
}
