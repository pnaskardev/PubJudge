package config

import (
	"fmt"
	"os"

	cache "github.com/pnaskardev/pubjudge/gateway/cache"
	mongo_db "github.com/pnaskardev/pubjudge/gateway/db"
	mongo_types "github.com/pnaskardev/pubjudge/gateway/types/mongo_types"
	redis_types "github.com/pnaskardev/pubjudge/gateway/types/redis_types"
)

type AppConfig struct {
	Mongo mongo_types.MongoClientConnectionParams
	Redis redis_types.RedisClientConnectionParams
}

type App struct {
	Db    *mongo_types.MongoClientStruct
	Cache *redis_types.RedisClientStruct
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
		Mongo: mongo_types.MongoClientConnectionParams{
			Username: getEnvOrPanic("MONGO_USERNAME"),
			Password: getEnvOrPanic("MONGO_PASSWORD"),
		},
		Redis: redis_types.RedisClientConnectionParams{
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

	mongo_connection_result, err := mongo_db.ConnectToMongoDB(&config.Mongo)
	if mongo_connection_result == nil || err != nil {
		panic("Mongo Connection Failed")
	}

	redis_connnection_result, err := cache.NewRedisClient(&config.Redis)
	if redis_connnection_result == nil || err != nil {
		panic("Redis Connection Failed")
	}

	return &App{Db: mongo_connection_result, Cache: redis_connnection_result}, nil

}

func CloseDBConnection(app *App) {
	mongo_db.CloseDB(app.Db.Client)
}
