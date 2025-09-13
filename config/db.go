package config

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDB() {
	uri := os.Getenv("MONGODB_CONNECTION_STRING")
	if uri == "" {
		logrus.Fatal("Mongo DB Connection is not set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Fatal("Failed to connect to MongoDB: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logrus.Fatal("MongoDB not responding: ", err)
	}

	mongoClient = client
	logrus.Info("Connected to MongoDB")
}
func GetMongoClient() *mongo.Client {
	return mongoClient
}

//func GetMongoDBConnectionString() string {
//	conn := os.Getenv("MONGODB_CONNECTION_STRING")
//	if conn == "" {
//		return "mongodb://localhost:27017"
//	}
//	return conn
//}
