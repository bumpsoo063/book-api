package database

import (
	"context"
	"os"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
}

func GetRedis() *redis.Client {
	return rdb
}

func SetToken(userId string, token *model.Token, redis *redis.Client, f bool) error {
	accessExpire := time.Unix(token.AccessExpire, 0)
	refreshExpire := time.Unix(token.RefreshExpire, 0)
	now := time.Now()
	if err := redis.Set(context.Background(), token.AccessUuid, userId, accessExpire.Sub(now)).Err(); err != nil {
		return err
	}
	if f == true {
		return nil
	}
	if err := redis.Set(context.Background(), token.RefreshUuid, userId, refreshExpire.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}
