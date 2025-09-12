package repository

import (
	"smart-edu-api/config"
	"smart-edu-api/data/model/response"
	"smart-edu-api/entity"
	"smart-edu-api/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateModel(model entity.Model) (entity.Model, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")

	model.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(helper.GetContext(), model)
	if err != nil {
		return entity.Model{}, err
	}

	return model, nil
}

func GetAllModel() ([]response.ModelResponse, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("models")

	var result entity.Model
	return result.GetAll(collection)
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
