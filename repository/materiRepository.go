package repository

import (
	"smart-edu-api/config"
	respond "smart-edu-api/data/baseMateri/response"
	"smart-edu-api/entity"
	"smart-edu-api/helper"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// CreateMateriPokok creates a new MateriPokok in the database
func CreateMateriPokok(baseMateri entity.MateriPokok) (entity.MateriPokok, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")
	baseMateri.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(helper.GetContext(), baseMateri)
	if err != nil {
		return entity.MateriPokok{}, err
	}
	return baseMateri, nil
}

// GetAllMateriPokok retrieves all MateriPokok from the database
func GetAllMateriPokok() ([]respond.GetMateriPokokResponse, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	var result entity.MateriPokok
	return result.GetAll(collection)
}

// GetMateriPokokByID retrieves a MateriPokok by its ID from the database
func GetMateriPokokByID(id string) (*entity.MateriPokok, error) {
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

// UpdateMateriPokok updates an existing MateriPokok in the database
func UpdateMateriPokok(materi *entity.MateriPokok) (*entity.MateriPokok, error) {
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

// DeleteMateriPokok deletes a MateriPokok from the database
func DeleteMateriPokok(materi *entity.MateriPokok) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	_, err := collection.DeleteOne(helper.GetContext(), bson.M{"_id": materi.ID})
	if err != nil {
		return err
	}
	return nil
}
