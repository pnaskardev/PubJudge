package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubmissionResponse struct {
	ID          primitive.ObjectID `json:"id"`
	UserID      primitive.ObjectID `json:"user_id"`
	Status      string             `json:"status"` // Always "Pending" for new submissions
	Language    string             `json:"language"`
	SubmittedAt time.Time          `json:"submitted_at"`
	Message     string             `json:"message"` // Optional: user-friendly text
}

func SubmissionSuccess(data *entities.Submission) *fiber.Map {

	response := SubmissionResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		Status:      "Pending",
		Language:    data.Language,
		SubmittedAt: data.SubmittedAt,
		// Message:     "Submission received and queued for evaluation.",
	}
	return &fiber.Map{
		"status": true,
		"data":   response,
		"error":  nil,
	}
}
