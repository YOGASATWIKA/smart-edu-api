package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
