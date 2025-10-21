package repository

import (
	"context"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

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

func GetEbookByModulId(id string) (*entity.Ebook, error) {
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

func UpdateEbookById(id string, ebook entity.Ebook) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("ebook")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":        ebook.Title,
			"modul":        ebook.ModuleId,
			"parts":        ebook.Parts,
			"html_content": ebook.HtmlContent,
			"json_content": ebook.JsonContent,
			"updated_at":   time.Now(),
		},
	}

	_, err = collection.UpdateOne(helper.GetContext(), bson.M{"modul": objectID}, update)
	return err
}

func CreateLog(ctx context.Context, ebook entity.Ebook) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("ebook_log")
	_, err := collection.InsertOne(ctx, ebook)
	if err != nil {
		return err
	}
	return nil
}

func GetEbookById(id string) (*entity.Ebook, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("ebook")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var ebook entity.Ebook
	err = collection.FindOne(helper.GetContext(), filter).Decode(&ebook)
	if err != nil {
		return nil, err
	}

	return &ebook, nil
}
