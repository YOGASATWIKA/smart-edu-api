package repository

import (
	"context"
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetOutlineByMateriPokokId(ctx context.Context, id string) (*entity.Outline, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("outline")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"materi_pokok._id": objectID}
	var materi entity.Outline
	err = collection.FindOne(ctx, filter).Decode(&materi)
	if err != nil {
		return nil, err
	}
	return &materi, nil
}

func UpdateOutline(outline *entity.Outline) (*entity.Outline, error) {

	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("outline")

	filter := bson.M{"_id": outline.ID}
	update := bson.M{"$set": outline}

	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(helper.GetContext(), filter, update, opts)
	if err != nil {
		return nil, err
	}
	return outline, nil
}

func CreateOutline(ctx context.Context, outline entity.Outline) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("outline")
	_, err := collection.InsertOne(ctx, outline)
	if err != nil {
		return err
	}

	return nil
}
