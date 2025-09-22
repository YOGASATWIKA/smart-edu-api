package modul

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetActivity struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Namajabatan string             `json:"nama_jabatan" bson:"nama_jabatan"`
	Status      string             `json:"status" bson:"status"`
	State       string             `json:"state" bson:"state"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
