package cache

import (
	"Caro_Game/common"
	"github.com/go-redis/redis/v8"
)

func AddToken(idUser int, token string, redis *redis.Client) {
	err := redis.Set(redis.Context(), common.IntToString(idUser), token, 0).Err()
	if err != nil {
		return
	}
}

func GetTokenRedis(idUser int, redis *redis.Client) string {
	result, err := redis.Get(redis.Context(), common.IntToString(idUser)).Result()
	if err != nil {
		return "Error"
	}
	return result
}

func DeleteTokenRedis(idUser int, redis *redis.Client) string {
	_, err := redis.Del(redis.Context(), common.IntToString(idUser)).Result()
	if err != nil {
		return "error"
	}
	return "deleted"
}
