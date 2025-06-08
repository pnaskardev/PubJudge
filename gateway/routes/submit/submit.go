package submit

import (
	"github.com/gofiber/fiber/v2"
)

type SubmissionRoutes struct {
	app *fiber.App
}

func NewSubmissionRoutes(app *fiber.App) *SubmissionRoutes {
	return &SubmissionRoutes{app: app}
}

func (r *SubmissionRoutes) Register() {
	// group := r.app.Group("/api/submission")
}
