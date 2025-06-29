package cache

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/pnaskardev/pubjudge/worker/types"
	"github.com/redis/go-redis/v9"
)

var (
	once        sync.Once
	redisClient *redis.Client
)

func NewRedisClient(params *types.RedisClientConnectionParams) (*types.RedisClientStruct, error) {
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

	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: params.Username,
			Password: params.Password,
			DB:       int(dbInt64),
		})
	})

	// Test connection
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &types.RedisClientStruct{
		Client: redisClient,
		DB:     int(dbInt64),
	}, nil
}

func CloseRedisConnection() {
	redisClient.Conn().Close()
}
