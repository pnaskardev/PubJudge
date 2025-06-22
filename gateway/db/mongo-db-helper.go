package mongo_db

import (
	"context"
	"fmt"
	"sync"

	mongo_types "github.com/pnaskardev/pubjudge/gateway/types/mongo_types"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const dbName = "pub-judge"

var (
	Once        sync.Once
	mongoClient *mongo.Client
	initErr     error
)

func ConnectToMongoDB(params *mongo_types.MongoClientConnectionParams) (*mongo_types.MongoClientStruct, error) {

	var uri string = "mongodb+srv://pnaskardev:<db_password>@stream.vuocv.mongodb.net/?retryWrites=true&w=majority&appName=stream"
	if params.Username != "" && params.Password != "" {
		uri = fmt.Sprintf("mongodb+srv://%s:%s@stream.vuocv.mongodb.net/?retryWrites=true&w=majority&appName=stream", params.Username, params.Password)
	}
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	// client, err := mongo.Connect(opts)
	// if err != nil {
	// 	panic(err)
	// }

	Once.Do(func() {
		mongoClient, initErr = mongo.Connect(opts)
		// error detected now get return and go to default error handler
		if initErr != nil {
			return
		}
		// Send a ping to confirm a successful connection
		if initErr = mongoClient.Ping(context.TODO(), readpref.Primary()); initErr != nil {
			return
		}
		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	})

	if initErr != nil {
		return nil, fmt.Errorf("MongoDB connection failed: %w", initErr)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	var database mongo.Database = *mongoClient.Database(dbName)

	return &mongo_types.MongoClientStruct{Client: mongoClient, Database: &database}, nil

}

func CloseDB() error {
	err := mongoClient.Disconnect(context.TODO())

	if err != nil {
		return err

	}
	return nil
}
