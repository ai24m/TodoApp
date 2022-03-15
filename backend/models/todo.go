package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	//public fields are captialized + mongodb stores data in BSON format
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Description string             `json:"status,omitempty" bson:"status,omitempty"`
	IsComplete  bool               `json:"is_complete,omitempty" bson:"is_complete,omitempty"`
}
