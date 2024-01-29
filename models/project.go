package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Project struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      string             `json:"userId"`
	Name        string             `json:"name"`
	Description string             `json:"string"`
	CreatedAt   time.Time          `json:"createdAt"`
}
