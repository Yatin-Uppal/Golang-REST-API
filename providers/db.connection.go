package providers

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
function to createDatabase connectivity
returns (mongoClientInstance, errorObject)
*/
func GetMongoDBClient() (*mongo.Client, error) {
	var mongoClientInstance *mongo.Client
	var mongoOnce sync.Once
	var clientInstanceError error

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		mongoClientInstance = client
		clientInstanceError = err
	})
	return mongoClientInstance, clientInstanceError
}
