package entity

import (
	"context"
	respond "smart-edu-api/data/model/response"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	Model            string             `json:"model" bson:"model"`
	Status           string             `json:"status" bson:"status"`
	PromtContext     string             `json:"promt_context" bson:"promt_context"`
	PromtInstruction string             `json:"promt_instruction" bson:"promt_instruction"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt         time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

func (b *Model) GetAll(collection *mongo.Collection) ([]respond.ModelResponse, error) {
	var results []respond.ModelResponse

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"status": bson.M{"$ne": "DELETED"}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
