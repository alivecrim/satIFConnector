package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Active      bool               `bson:"active" json:"active"`
}
