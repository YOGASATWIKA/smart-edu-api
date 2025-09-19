package helper

import (
	"context"
	"smart-edu-api/config"
	"smart-edu-api/embeded"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func IsNameExists(Name string) (bool, error) {
	client := config.GetMongoClient()
	collection := client.Database("smart_edu").Collection("skb")
	filter := bson.M{"nama_jabatan": Name, "status": bson.M{"$ne": "DELETED"}}

	var result embeded.MateriPokok
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil // nama belum ada
		}
		return false, err // error lain
	}
	return true, nil // nama sudah ada
}
