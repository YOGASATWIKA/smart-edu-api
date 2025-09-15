package entity

import (
	"smart-edu-api/embeded"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Outline struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
	MateriPokok embeded.MateriPokokEmbed `json:"materi_pokok" bson:"materi_pokok"`
	Outline     embeded.Outline          `json:"outline,omitempty" bson:"outline,omitempty"`
	CreatedAt   time.Time                `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt    time.Time                `json:"delete_at" bson:"delete_at,omitempty"`
}
