package cache

import (
	"Caro_Game/common"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	TIME_EXPIRE = 60 // 60s het han
)

func AddToken(idUser int, token string, redis *redis.Client) {
	err := redis.Set(redis.Context(), token, common.IntToString(idUser), time.Hour).Err()
	//redis.Expire(redis.Context(), common.IntToString(idUser), TIME_EXPIRE)
	if err != nil {
		return
	}
}

func GetTokenRedis(token string, redis *redis.Client) string {
	result, err := redis.Get(redis.Context(), token).Result()
	if err != nil {
		return "Error"
	}
	return result
}

func DeleteTokenRedis(token string, redis *redis.Client) string {
	_, err := redis.Del(redis.Context(), token).Result()
	if err != nil {
		return "error"
	}
	return "deleted"
}
