package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamAchievement struct {
	Id 					primitive.ObjectID		`json:"id" bson:"id,omitempty"`
	TeamId				primitive.ObjectID		`json:"teamId" bson:"teamId,omitempty"`
	AchievementId		primitive.ObjectID		`json:"achievementId" bson:"achievementId,omitempty"`		
	CreatedAt			time.Time				`json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt			time.Time				`json:"updatedAt" bson:"updatedAt,omitempty"`
}