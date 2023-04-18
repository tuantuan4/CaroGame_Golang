package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

const (
	ADDRESS  = "localhost:6379"
	PASSWORD = ""
	DATABASE = 0
)

func CreateRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     ADDRESS,
		Password: PASSWORD,
		DB:       DATABASE,
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
