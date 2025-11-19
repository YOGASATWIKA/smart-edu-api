package repository

import (
	"smart-edu-api/config"
	"smart-edu-api/entity"
	"smart-edu-api/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(modul *entity.User) (*entity.User, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")
	filter := bson.M{"_id": modul.ID}
	update := bson.M{"$set": modul}
	_, err := collection.UpdateOne(helper.GetContext(), filter, update)
	if err != nil {
		return modul, err
	}
	return modul, nil
}

func GetUserById(id string) (*entity.User, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var user entity.User
	err = collection.FindOne(helper.GetContext(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*entity.User, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("users")

	filter := bson.M{"email": email}
	var user entity.User
	err := collection.FindOne(helper.GetContext(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
