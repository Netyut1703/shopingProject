package store

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Database
}

func New() *Store {

	s := &Store{}
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	s.db = client.Database("shopingProject")
	return s
}
