package redis

import (
	"fmt"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	Client *redis.Client
)

func Setup() {
	Client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	fmt.Println("redis addr:", viper.GetString("redis.addr"))

	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func TestSetup() {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	// 打开下面命令可以测试链接关闭的情况
	// defer mr.Close()

	Client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})
	fmt.Println("mini redis addr:", mr.Addr())
}
