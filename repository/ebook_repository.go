package repository

import (
	"context"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMateri(ctx context.Context, ebook entity.Ebook) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("ebook")
	_, err := collection.InsertOne(ctx, ebook)
	if err != nil {
		return err
	}
	return nil
}

func GetEbookByModulId(ctx context.Context, id string) (*entity.Ebook, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("ebook")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"modul": objectID}
	var materi entity.Ebook
	err = collection.FindOne(helper.GetContext(), filter).Decode(&materi)
	if err != nil {
		return nil, err
	}
	return &materi, nil
}

//
//func UpdateMateri(ebook *entity.Ebook) (*entity.Ebook, error) {
//
//	client := config.GetMongoClient()
//	collection := client.Database("smart_edu").Collection("materi")
//
//	filter := bson.M{"_id": ebook.ID}
//	update := bson.M{"$set": ebook}
//
//	opts := options.Update().SetUpsert(true)
//	_, err := collection.UpdateOne(helper.GetContext(), filter, update, opts)
//	if err != nil {
//		return nil, err
//	}
//	return ebook, nil
//}

func CreateLog(ctx context.Context, ebook entity.Ebook) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("ebook_log")
	_, err := collection.InsertOne(ctx, ebook)
	if err != nil {
		return err
	}
	return nil
}
