package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Model       string             `json:"model" bson:"model"`
	Description string             `bson:"description" json:"description"`
	Steps       []PromptStep       `bson:"steps" json:"steps"`
	Variables   []string           `bson:"variables" json:"variables"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt    time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

type PromptStep struct {
	Role    string `bson:"role" json:"role"`
	Content string `bson:"content" json:"content"`
}
