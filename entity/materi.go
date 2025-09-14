package entity

import (
	"smart-edu-api/embeded"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Materi struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Namajabatan  string             `json:"nama_jabatan" bson:"nama_jabatan"`
	Tugasjabatan []string           `json:"tugas_jabatan" bson:"tugas_jabatan"`
	Keterampilan []string           `json:"keterampilan" bson:"keterampilan"`
	Klasifikasi  string             `json:"klasifikasi" bson:"klasifikasi"`
	Status       string             `json:"status" bson:"status"`
	Stage        string             `json:"stage" bson:"stage"`
	Outline      embeded.Outline    `json:"outline" bson:"outline"`
	Ebook        embeded.Ebook      `json:"ebook" bson:"ebook"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt     time.Time          `json:"delete_at" bson:"delete_at,omitempty"`
}

//func (b *Materi) GetAll(collection *mongo.Collection) ([]respond.GetMateriPokokResponse, error) {
//	var results []respond.GetMateriPokokResponse
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	filter := bson.M{"status": bson.M{"$ne": "DELETED"}}
//	cursor, err := collection.Find(ctx, filter)
//	if err != nil {
//		return nil, err
//	}
//	defer cursor.Close(ctx)
//
//	if err := cursor.All(ctx, &results); err != nil {
//		return nil, err
//	}
//
//	return results, nil
//}
