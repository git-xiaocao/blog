package rdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"server/config"
)

var Client *redis.Client

func init() {
	rConfig := config.Config().Redis
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rConfig.Host, rConfig.Port),
		Password: rConfig.Password,
		DB:       rConfig.DB,
	})

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
