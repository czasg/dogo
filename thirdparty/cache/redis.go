package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"proj/lifecycle"
)

func NewRedis(ctx context.Context, cfg lifecycle.RedisConfig) (*redis.Client, error) {
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
