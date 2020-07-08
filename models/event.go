package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Event Model
type Event struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Date        string             `json:"date" bson:"date,omitempty"`
	Zones       []Zone             `json:"zones" bson:"zones,omitempty"`
}
