package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Ctx = context.Background()
)

func Perform() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("%s: %v", "Redis Ping Failed", err)
	}

	return rdb
}
