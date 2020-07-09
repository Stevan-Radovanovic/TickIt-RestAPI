package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Order Model
type Order struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Ticket string             `json:"ticket" bson:"ticket,omitempty"`
	Date   string             `json:"date" bson:"date,omitempty"`
	Amount string             `json:"amount" bson:"amount,omitempty"`
}
