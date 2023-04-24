package cache

import (
	"Caro_Game/common"
	"Caro_Game/models"
	"context"
	"encoding/json"
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

func AddUserRedis(token string, user models.User, redis *redis.Client) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return
	}
	err = redis.HSet(context.Background(), token, userJSON).Err()
	if err != nil {
		return
	}
}

func GetRoleUserRedis(token string, redis *redis.Client) string {
	userJSON, err1 := redis.HGet(context.Background(), token, "").Result()
	if err1 != nil {
		return "Error convert JSON"
	}
	var user models.User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		return "Error"
	}
	return common.IntToString(int(user.RoleID))
}
