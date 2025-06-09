package redis_types

import (
	"github.com/redis/go-redis/v9"
)

// Params to configure a Redis client
type RedisClientConnectionParams struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// Wrapper around redis.Client
type RedisClientStruct struct {
	Client *redis.Client
	DB     int
}
