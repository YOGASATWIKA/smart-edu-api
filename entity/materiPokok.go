package entity

import (
	"context"
	"time"

	respond "smart-edu-api/data/baseMateri/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MateriPokok struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Namajabatan  string             `json:"nama_jabatan" bson:"nama_jabatan"`
	Tugasjabatan []string           `json:"tugas_jabatan" bson:"tugas_jabatan"`
	Keterampilan []string           `json:"keterampilan" bson:"keterampilan"`
	Klasifikasi  string             `json:"klasifikasi" bson:"klasifikasi"`
	Status       string             `json:"status" bson:"status"`
	Outline      Outline            `json:"outline" bson:"outline"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt     time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

type Outline struct {
	ListMateri []Materi `json:"list_materi" bson:"list_materi"`
}
type Materi struct {
	MateriPokok   string      `json:"materi_pokok" bson:"materi_pokok"`
	ListSubMateri []SubMateri `json:"list_sub_materi" bson:"list_sub_materi"`
}
type SubMateri struct {
	SubMateriPokok string   `json:"sub_materi_pokok" bson:"sub_materi_pokok"`
	ListMateri     []string `json:"list_materi" bson:"list_materi"`
}

func (b *MateriPokok) GetAll(collection *mongo.Collection) ([]respond.GetMateriPokokResponse, error) {
	var results []respond.GetMateriPokokResponse

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
