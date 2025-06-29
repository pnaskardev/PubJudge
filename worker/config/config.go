package config

import (
	"fmt"
	"os"

	"github.com/pnaskardev/pubjudge/worker/cache"
	"github.com/pnaskardev/pubjudge/worker/types"
)

type AppConfig struct {
	Mongo types.MongoClientConnectionParams
	Redis types.RedisClientConnectionParams
}

type App struct {
	Db    *types.MongoClientStruct
	Cache *types.RedisClientStruct
}

func getEnvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s is not set", key))
	}
	return val
}

func loadConfig() *AppConfig {
	return &AppConfig{
		Mongo: types.MongoClientConnectionParams{
			Username: getEnvOrPanic("MONGO_USERNAME"),
			Password: getEnvOrPanic("MONGO_PASSWORD"),
		},
		Redis: types.RedisClientConnectionParams{
			Username: getEnvOrPanic("REDIS_USERNAME"),
			Password: getEnvOrPanic("REDIS_PASSWORD"),
			Host:     getEnvOrPanic("REDIS_HOST"),
			Port:     getEnvOrPanic("REDIS_PORT"),
			Database: getEnvOrPanic("REDIS_DATABASE"),
		},
	}
}

func Init() (*App, error) {

	config := loadConfig()

	// mongo_connection_result, err := mongo_db.ConnectToMongoDB(&config.Mongo)
	// if mongo_connection_result == nil || err != nil {
	// 	panic("Mongo Connection Failed")
	// }

	redis_connnection_result, err := cache.NewRedisClient(&config.Redis)
	if redis_connnection_result == nil || err != nil {
		panic("Redis Connection Failed")
	}

	return &App{Db: nil, Cache: redis_connnection_result}, nil

}

func CloseCacheConnection() {
	cache.CloseRedisConnection()
}
