package cache

import (
	"context"
	"fmt"
	"strconv"

	redis_types "github.com/pnaskardev/pubjudge/gateway/types/redis_types"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(params *redis_types.RedisClientConnectionParams) (*redis_types.RedisClientStruct, error) {
	addr := fmt.Sprintf("%s:%s", params.Host, params.Port)

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     addr,
	// 	Username: params.Username,
	// 	Password: params.Password,
	// 	DB:       params.Database,
	// })

	dbInt64, err := strconv.ParseInt(params.Database, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid database number: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: params.Username,
		Password: params.Password,
		DB:       int(dbInt64),
	})

	// Test connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &redis_types.RedisClientStruct{
		Client: rdb,
		DB:     int(dbInt64),
	}, nil
}
