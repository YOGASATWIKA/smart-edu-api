package repository

import (
	"context"
	"smart-edu-api/config"
	modul "smart-edu-api/data/modul/response"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateModul(modul entity.Modul) (entity.Modul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")
	modul.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(helper.GetContext(), modul)
	if err != nil {
		return entity.Modul{}, err
	}
	return modul, nil
}

func UpdateModul(ctx context.Context, modul entity.Modul) (entity.Modul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")
	filter := bson.M{"_id": modul.ID}
	update := bson.M{"$set": modul}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return modul, err
	}
	return modul, nil
}

func GetAllModul() ([]modul.GetAllModul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")

	var results []modul.GetAllModul

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"status": bson.M{"$ne": "DELETED"}, "state": bson.M{"$ne": "EBOOK"}}
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

func GetAllEbook() ([]modul.GetAllModul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")

	var results []modul.GetAllModul

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"status": bson.M{"$ne": "DELETED"},
		"state":  bson.M{"$nin": []string{"DRAFT", "OUTLINE"}},
	}
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

func GetModulById(id string) (entity.Modul, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("modul")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Modul{}, err
	}
	filter := bson.M{"_id": objectID}
	var modul entity.Modul
	err = collection.FindOne(helper.GetContext(), filter).Decode(&modul)
	if err != nil {
		return modul, err
	}
	return modul, nil
}
