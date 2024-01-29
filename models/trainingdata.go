package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrainingData struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	ProjectId string             `json:"projectId"`
	BotId     string             `json:"botId"`
	Files     []string           `json:"files"`
}
