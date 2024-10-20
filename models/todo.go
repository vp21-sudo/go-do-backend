package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title" validate:"required,min=3,max=100"`
	Completed bool               `json:"completed" bson:"completed"`
}
