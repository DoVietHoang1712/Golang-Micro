package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieGenre struct {
	ID          primitive.ObjectID `bson:"id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}
