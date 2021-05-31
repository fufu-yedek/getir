package mongo

import (
	"context"
	"github.com/fufuceng/getir-challange/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

//Initialize responsible to initialize mongo db connection
func Initialize(config config.Mongo) error {
	opts := options.Client().ApplyURI(config.Uri)
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		return err
	}

	db = client.Database(config.Name)
	return nil
}

//DB responsible to serve mongo connection object
func DB() *mongo.Database {
	return db
}
