package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User Constructs your User model under entities.
type User struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Firstname string             `json:"first_name" bson:"first_name"`
	Lastname  string             `json:"last_name" bson:"last_name,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `bson:"password"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Books
type DeleteRequest struct {
	ID string `json:"id"`
}
