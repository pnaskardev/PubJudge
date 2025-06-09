package config

import (
	"os"

	"github.com/pnaskardev/pubjudge/gateway/constants/types"
	mongo_db "github.com/pnaskardev/pubjudge/gateway/db"
)

func Init() {
	mongo_username := os.Getenv("MONGO_USERNAME")
	mongo_password := os.Getenv("MONGO_PASSWORD")
	if mongo_username == "" || mongo_password == "" {
		panic("Mongo DB Connection Parameters not found")
	}

	mongo_db_params := types.MongoClientConnectionParams{
		Username: mongo_username,
		Password: mongo_password,
	}

	mongo_connection_result := mongo_db.ConnectToMongoDB(&mongo_db_params)
	if mongo_connection_result == nil {
		panic("Mongo Connection Failed")
	}
}
