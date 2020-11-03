package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person Model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty" validation:"required"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validation:"required"`
}
