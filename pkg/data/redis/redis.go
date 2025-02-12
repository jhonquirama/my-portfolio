package redis

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
	customLogger "github.com/jhonquirama/hexa-scaffolding-ms/pkg/log"
	apm "github.com/jhonquirama/hexa-scaffolding-ms/pkg/monitor/elastic-apm"
)

var (
	RedisNil = redis.Nil // nolint: gochecknoglobals
)

type (
	Cache interface {
		Set(ctx context.Context, key, value string, ttl time.Duration) error
		Get(ctx context.Context, key string) (string, error)
		SetNX(ctx context.Context, key string, body any, timeout time.Duration) (bool, error)
	}

	Config interface {
		Addr() string
		Password() string
		DB() int
		PoolSize() int
		Timeout() int64
	}

	client struct {
		redis *redis.Client
	}
)

func NewRedisClient(ctx context.Context, config Config) (Cache, error) {
	var (
		client  client
		timeout = time.Duration(config.Timeout()) * time.Second
		err     error
	)

	client.redis = redis.NewClient(&redis.Options{
		Addr: config.Addr(),
		OnConnect: func(context.Context, *redis.Conn) error {
			customLogger.Info(ctx, "Connected to redis...")
			return nil
		},
		Password:     config.Password(),
		DB:           config.DB(),
		DialTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		PoolSize:     config.PoolSize(),
		PoolTimeout:  timeout,
	})

	client.redis.AddHook(apm.NewRedisHook())

	if err = client.redis.Ping(context.TODO()).Err(); err != nil {
		return nil, err
	}

	return &client, err
}

func (c *client) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.redis.Set(ctx, key, value, ttl).Err()
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	return c.redis.Get(ctx, key).Result()
}

func (c *client) SetNX(ctx context.Context, key string, body any, timeout time.Duration) (bool, error) {
	return c.redis.SetNX(ctx, key, body, timeout).Result()
}
