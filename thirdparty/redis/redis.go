package redis

import (
	"context"
	"proj/public/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, config.RedisConfig) (*redis.Client, error) {
	ins := redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MaxRetries:   cfg.MaxRetries,
		MinIdleConns: cfg.MinIdleSize,
	})
	err := ins.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return ins, nil
}
