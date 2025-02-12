package apm

import (
	redis "github.com/go-redis/redis/v8"
	apmRedis "go.elastic.co/apm/module/apmgoredisv8/v2"
)

func NewRedisHook() redis.Hook {
	return apmRedis.NewHook()
}
