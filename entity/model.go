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
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Model     string             `json:"model" bson:"model"`
	Promt     Promt              `json:"promt" bson:"promt"`
	Type      string             `json:"type" bson:"type"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt  time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

type Promt struct {
	SystemPrompt string   `json:"system_prompt" bson:"system_prompt"`
	UserPrompts  []string `json:"user_prompts" bson:"user_prompts"`
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
