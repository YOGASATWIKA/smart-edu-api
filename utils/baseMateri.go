package utils

import (
	"context"
	"smart-edu-api/config"
	respond "smart-edu-api/data/baseMateri/response"
	"smart-edu-api/model"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetBaseMateri() ([]respond.BaseMateriRespond, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	var result model.BaseMateri
	return result.GetAll(collection)
}

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func CreateBaseMateri(baseMateri model.BaseMateri) (model.BaseMateri, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	// Set the ID of the base materi to the inserted ID
	baseMateri.ID = primitive.NewObjectID()

	// Insert the base materi into the collection
	_, err := collection.InsertOne(GetContext(), baseMateri)
	if err != nil {
		return model.BaseMateri{}, err
	}


	return baseMateri, nil
}
func GetCurrentTime() time.Time {
    return time.Now().UTC()
}

func IsJobIDExists(Jobid int64) (bool, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")
	filter := bson.M{"job_id": Jobid, "status": bson.M{"$ne": "DELETED"}}

	var result model.BaseMateri
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil // jobid belum ada
		}
		return false, err // error lain
	}
	return true, nil // jobid sudah ada
}

func IsNameExists(Name string) (bool, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")
	filter := bson.M{"nama_jabatan": Name, "status": bson.M{"$ne": "DELETED"}}

	var result model.BaseMateri
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil // nama belum ada
		}
		return false, err // error lain
	}
	return true, nil // nama sudah ada
}

func GetBaseMateriByID(id string) (*model.BaseMateri, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var materi model.BaseMateri
	err = collection.FindOne(GetContext(), filter).Decode(&materi)
	if err != nil {
		return nil, err
	}
	return &materi, nil
}

func UpdateBaseMateri(materi *model.BaseMateri) (*model.BaseMateri, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	filter := bson.M{"_id": materi.ID}
	update := bson.M{"$set": materi}

	_, err := collection.UpdateOne(GetContext(), filter, update)
	if err != nil {
		return nil, err
	}
	return materi, nil
}

// Fungsi untuk menghapus base materi
func DeleteBaseMateri(materi *model.BaseMateri) error {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")

	_, err := collection.DeleteOne(GetContext(), bson.M{"_id": materi.ID})
	if err != nil {
		return err
	}
	return nil
}




