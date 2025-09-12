package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	GoogleID string             `bson:"googleId,omitempty" json:"googleId,omitempty"`
	Email    string             `bson:"email,omitempty" json:"email,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Picture  string             `bson:"picture,omitempty" json:"picture,omitempty"`
}
