package model

import (
	// "context"
	"context"
	"time"

	respond "smart-edu-api/data/baseMateri/response"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseMateri struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Jobid           int64          `json:"job_id" bson:"job_id"`
	Namajabatan     string             `json:"nama_jabatan" bson:"nama_jabatan"`
	Tugasjabatan    []string             `json:"tugas_jabatan" bson:"tugas_jabatan"`
	Keterampilan    []string             `json:"keterampilan" bson:"keterampilan"`
	Klasifikasi     string             `json:"klasifikasi" bson:"klasifikasi"`
	Status 			string             `json:"status" bson:"status"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeleteAt       	time.Time        `json:"delete_at" bson:"delete_at,omitempty"`
}

func (b *BaseMateri) GetAll(collection *mongo.Collection) ([]respond.BaseMateriRespond, error) {
	var results []respond.BaseMateriRespond

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
