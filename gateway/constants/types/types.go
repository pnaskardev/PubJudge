package types

import "go.mongodb.org/mongo-driver/v2/mongo"

type MongoClientConnectionParams struct {
	Username string
	Password string
}

type MongoClientStruct struct {
	Client   *mongo.Client
	Database *mongo.Database
}