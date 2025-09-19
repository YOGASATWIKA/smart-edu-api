package repository

import (
	"context"
	"smart-edu-api/config"
	"smart-edu-api/data/modul/response"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateModul(baseMateri entity.Modul) (entity.Modul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")
	baseMateri.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(helper.GetContext(), baseMateri)
	if err != nil {
		return entity.Modul{}, err
	}
	return baseMateri, nil
}

func GenerateOutline(ctx context.Context, modul entity.Modul) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")
	_, err := collection.InsertOne(ctx, modul)
	if err != nil {
		return err
	}

	return nil
}

func GetAllModul() ([]response.GetAllModul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")

	var results []response.GetAllModul

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

//
//func GetOutlineByMateriPokokId(ctx context.Context, id string) (*entity.Outline, error) {
//	client := config.GetMongoClient()
//	collection := client.Database("smart_edu").Collection("outline")
//
//	objectID, err := primitive.ObjectIDFromHex(id)
//	if err != nil {
//		return nil, err
//	}
//
//	filter := bson.M{"materi_pokok._id": objectID}
//	var materi entity.Outline
//	err = collection.FindOne(ctx, filter).Decode(&materi)
//	if err != nil {
//		return nil, err
//	}
//	return &materi, nil
//}
