package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Gender string             `bson:"gender" json:"gender"`
	Age    int                `bson:"age" json:"age"`
}