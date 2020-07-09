package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Zone Model
type Zone struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Amount string             `json:"amount" bson:"amount,omitempty"`
}
