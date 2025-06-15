package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname,omitempty"` // optional
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

// User Constructs your User model under entities.
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname"` // Make optional in JSON if needed
	Lastname  string             `json:"lastname,omitempty" bson:"lastname"`   // Already optional
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"-" bson:"password"` // Hide password in JSON responses
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Books
type DeleteRequest struct {
	ID string `json:"id"`
}
