package response

import (
	"smart-edu-api/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAllModul struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	MateriPokok entity.MateriPokok `json:"materi_pokok" bson:"materi_pokok"`
	Outline     entity.Outline     `json:"outline,omitempty" bson:"outline,omitempty"`
	Status      string             `json:"model" bson:"model"`
	State       string             `json:"model" bson:"model"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
