package initialize

import (
	"nocake/global"

	"github.com/go-redis/redis/v8"
)

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123qwe", //
		DB:       0,        // use default DB
	})
	global.Rdb = rdb
}
