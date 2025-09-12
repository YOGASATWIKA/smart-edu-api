package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ModelCopy struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Model     string             `json:"model" bson:"model"`
	Status    string             `json:"status" bson:"status"`
	Outline   OutlinePromt       `json:"outline" bson:"outline"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt  time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

type OutlinePromt struct {
	//Mengatur Konteks, Meningkatkan Kualitas dan Relevansi, Meminimalkan Ambiguitas
	RolePromt         string `json:"role_promt" bson:"role_promt"`
	VariableInjection string `json:"variable_injection" bson:"variable_injection"`
}
