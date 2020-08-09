package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
	Nil    = redis.Nil
)

func Setup() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	fmt.Println("redis addr:", viper.GetString("redis.addr"))

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	Client = rdb

	return rdb
}
