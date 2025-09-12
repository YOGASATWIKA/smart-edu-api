package response

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMateriPokokResponse struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Namajabatan  string             `json:"nama_jabatan" bson:"nama_jabatan"`
	Tugasjabatan []string           `json:"tugas_jabatan" bson:"tugas_jabatan"`
	Keterampilan []string           `json:"keterampilan" bson:"keterampilan"`
	Klasifikasi  string             `json:"klasifikasi" bson:"klasifikasi"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
