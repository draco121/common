package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Email     string             `json:"email"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Password  string             `json:"password"`
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Role      Role               `json:"role"`
}
