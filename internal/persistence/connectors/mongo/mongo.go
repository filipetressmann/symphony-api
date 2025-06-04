package mongo

import (
	"context"
	"fmt"
	"log"
	"symphony-api/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client *mongo.Client
}
// NewMongoConnection creates a new MongoConnection instance.
// It initializes the MongoDB client with the connection string
// constructed from environment variables for username and password.
// The connection string is in the format: mongodb://<username>:<password>@mongo:27017
// If the connection fails, it logs the error and exits the application.
// The MongoConnection struct holds the MongoDB client which can be used
// to interact with the MongoDB database.
func NewMongoConnection() *MongoConnection {
	username := config.GetEnv("MONGO_INITDB_ROOT_USERNAME", "root")
	password := config.GetEnv("MONGO_INITDB_ROOT_PASSWORD", "rootpassword")
	mongoURI := fmt.Sprintf("mongodb://%s:%s@mongo:27017", username, password)

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	return &MongoConnection {
		client: client,
	}
}
