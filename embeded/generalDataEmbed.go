package embeded

import "go.mongodb.org/mongo-driver/bson/primitive"

type MateriPokokEmbed struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Namajabatan string             `json:"nama_jabatan" bson:"nama_jabatan"`
}
