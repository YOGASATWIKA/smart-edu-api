package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Conn *mongo.Client
}

var db *Database

func New(connection string) (*Database, error) {
	if db != nil {
		return db, nil
	}

	opt := options.Client().ApplyURI(connection)

	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	def := &Database{
		Conn: client,
	}
	return def, nil
}
