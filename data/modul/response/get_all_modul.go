package modul

import (
	"smart-edu-api/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAllModul struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	MateriPokok entity.MateriPokok `json:"materi_pokok" bson:"materi_pokok"`
	Outline     entity.Outline     `json:"outline,omitempty" bson:"outline,omitempty"`
	Status      string             `json:"status" bson:"status"`
	State       string             `json:"state" bson:"state"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
