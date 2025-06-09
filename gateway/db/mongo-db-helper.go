package mongo_db

import (
	"context"
	"fmt"

	"github.com/pnaskardev/pubjudge/gateway/constants/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const dbName = "pub-judge"

func ConnectToMongoDB(params *types.MongoClientConnectionParams) *types.MongoClientStruct {

	var uri string = "mongodb+srv://pnaskardev:<db_password>@stream.vuocv.mongodb.net/?retryWrites=true&w=majority&appName=stream"
	if params.Username != "" && params.Password != "" {
		uri = fmt.Sprintf("mongodb+srv://%s:%s@stream.vuocv.mongodb.net/?retryWrites=true&w=majority&appName=stream", params.Username, params.Password)
	}
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	var database mongo.Database = *client.Database(dbName)

	return &types.MongoClientStruct{Client: client, Database: &database}

}
