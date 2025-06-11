package presenter

import (
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the presenter object which will be passed in the response by Handler
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname string             `json:"firstname"`
}

type UserWithSubmission struct {
	ID          primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	Firstname   string                `json:"firstname"`
	Submissions []entities.Submission `bson:"submissions" json:"submissions"`
}

// BookSuccessResponse is the singular SuccessResponse that will be passed in the response by
// Handler
func UserSuccessResponse(data *entities.User) *fiber.Map {
	book := User{
		ID: data.ID,
	}
	return &fiber.Map{
		"status": true,
		"data":   book,
		"error":  nil,
	}
}

// BooksSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func UsersSuccessResponse(data *[]User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

// BookErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
