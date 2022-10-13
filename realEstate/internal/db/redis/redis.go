package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	logg "realEstate/pkg/log"
)

var ctx = context.Background()

// initialize redis db
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		logg.Info("Can not connect to redis")

	}
	logg.Info("Connect to redis succesfully")
	logg.Info(pong)
	return rdb
}
