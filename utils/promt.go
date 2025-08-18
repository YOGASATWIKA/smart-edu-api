package utils

import (
	"smart-edu-api/config"
	"smart-edu-api/data/promt/response"
	"smart-edu-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreatePromt(promt model.Promt) (model.Promt, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("promts")

	promt.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(GetContext(), promt)
	if err != nil {
		return model.Promt{}, err
	}

	return promt, nil
}

func GetPromt() ([]response.PromtResponse, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("promts")

	var result model.Promt
	return result.GetAll(collection)
}

func GetPromtById(id string) (*model.Promt, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("promts")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var promt model.Promt
	err = collection.FindOne(GetContext(), filter).Decode(&promt)
	if err != nil {
		return nil, err
	}
	return &promt, nil
}

func UpdatePromt(promt *model.Promt) (*model.Promt, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("promts")

	filter := bson.M{"_id": promt.ID}
	update := bson.M{"$set": promt}

	_, err := collection.UpdateOne(GetContext(), filter, update)
	if err != nil {
		return nil, err
	}
	return promt, nil
}