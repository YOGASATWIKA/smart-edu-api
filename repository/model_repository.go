package repository

import (
	"context"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateModel(model entity.Model) (entity.Model, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")
	_, err := collection.InsertOne(helper.GetContext(), model)
	if err != nil {
		return entity.Model{}, err
	}

	return model, nil
}

func GetOutlineModel(typeParam string) ([]entity.Model, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")

	var results []entity.Model
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"status": bson.M{"$ne": "DELETED"},
	}
	switch typeParam {
	case "DRAFT", "OUTLINE", "EBOOK":
		filter["type"] = typeParam
	case "ALL":
		filter["type"] = bson.M{"$in": []string{"OUTLINE", "EBOOK"}}
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"updated_at": -1})

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func GetModelById(id string) (*entity.Model, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var model entity.Model
	err = collection.FindOne(helper.GetContext(), filter).Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func GetModelByModel(modelRequest string) (*entity.Model, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")

	filter := bson.M{"model": modelRequest}
	var model entity.Model
	err := collection.FindOne(helper.GetContext(), filter).Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func UpdateModel(model *entity.Model) (*entity.Model, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")

	filter := bson.M{"_id": model.ID}
	update := bson.M{"$set": model}

	_, err := collection.UpdateOne(helper.GetContext(), filter, update)
	if err != nil {
		return nil, err
	}
	return model, nil
}
