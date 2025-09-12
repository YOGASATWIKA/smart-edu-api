package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
)

func GetOutlineById(id string) (*entity.MateriPokok, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var materi entity.MateriPokok
	err = collection.FindOne(helper.GetContext(), filter).Decode(&materi)
	if err != nil {
		return nil, err
	}
	return &materi, nil
}
