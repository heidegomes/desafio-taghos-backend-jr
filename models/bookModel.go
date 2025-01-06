package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"titule"`
	Category string             `json:"category" bson:"category"`
	Author     string             `json:"author" bson:"author"`
	Synopsis   string             `json:"synopsis" bson:"synopsis"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}