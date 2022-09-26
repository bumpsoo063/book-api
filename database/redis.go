package database

import (
	"context"
	"os"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

var Redis *redis.Client

func ConnectRedis() *redis.Client {
	if Redis == nil {
		rdb := redis.NewClient(&redis.Options{
			Addr: os.Getenv("REDIS_ADDR"),
		})
		defer rdb.Close()
	}
	return Redis
}

func SetToken(user uuid.UUID, token *model.Token, redis *redis.Client) error {
	accessExpire := time.Unix(token.AccessExpire, 0)
	refreshExpire := time.Unix(token.RefreshExpire, 0)
	now := time.Now()
	if err := redis.Set(context.Background(), token.AccessToken, user.String(), accessExpire.Sub(now)).Err(); err != nil {
		return err
	}
	if err := redis.Set(context.Background(), token.RefreshToken, user.String(), refreshExpire.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}
