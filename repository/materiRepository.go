package repository

import (
	"context"
	"smart-edu-api/config"
	respond "smart-edu-api/data/baseMateri/response"
	"smart-edu-api/entity"
	"smart-edu-api/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

//========================================================MATERI POKOK=======================================================================

// CreateMateriPokok creates a new MateriPokok in the database
func CreateMateriPokok(baseMateri entity.Materi) (entity.Materi, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")
	baseMateri.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(helper.GetContext(), baseMateri)
	if err != nil {
		return entity.Materi{}, err
	}
	return baseMateri, nil
}

// GetAllMateriPokok retrieves all MateriPokok from the database
func GetAllMateriPokok() ([]respond.GetMateriPokokResponse, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	var results []respond.GetMateriPokokResponse

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

// GetMateriPokokByID retrieves a MateriPokok by its ID from the database
func GetMateriPokokByID(id string) (*entity.Materi, error) {
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

//========================================================OUTLINE

//========================================================FULL MATERI

//========================================================GENERAL

func DeleteMateri(materi *entity.Materi) (*entity.Materi, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	filter := bson.M{"_id": materi.ID}
	update := bson.M{"$set": materi}

	_, err := collection.UpdateOne(helper.GetContext(), filter, update)
	if err != nil {
		return nil, err
	}
	return materi, nil
}
