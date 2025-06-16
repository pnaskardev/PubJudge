package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Code        string             `bson:"code" json:"code"`
	Language    string             `bson:"language" json:"language"`
	SubmittedAt time.Time          `bson:"submitted_at" json:"submitted_at"`
}
