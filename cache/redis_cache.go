package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func CreateRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
func NewClientRedis() *redis.Client {
	rdb, err := CreateRedisClient()
	if err != nil {
		fmt.Println("failed to connect redis")
	} else {
		fmt.Println("connect redis successful")
	}

	return rdb
}
