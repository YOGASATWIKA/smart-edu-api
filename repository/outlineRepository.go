package repository

import (
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetOutlineById(id string) (*entity.Materi, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var materi entity.Materi
	err = collection.FindOne(helper.GetContext(), filter).Decode(&materi)
	if err != nil {
		return nil, err
	}
	return &materi, nil
}
