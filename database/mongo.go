package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(mongoUri string, databaseName string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	} else {
		db := client.Database(databaseName)
		return db
	}
}

func NewMongoDatabaseDefaults() *mongo.Database {
	var mongouri string = os.Getenv("MONGODB_URI")
	if mongouri == "" {
		mongouri = "mongodb://localhost:27017"
	}
	var mongodbname string = os.Getenv("MONGODB_DBNAME")
	if mongodbname == "" {
		mongodbname = "test-db"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongouri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	} else {
		db := client.Database(mongodbname)
		return db
	}
}
