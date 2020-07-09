package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//User Model
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}
