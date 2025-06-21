package submission_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/pnaskardev/pubjudge/gateway/pkg/submit"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleSubmit(service submit.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input entities.CreateSubmissionInput

		userIdObjectId := c.Locals("user_id").(primitive.ObjectID)

		if err := c.BodyParser(&input); err != nil {
			return presenter.BadRequest(c, "Inavlid Payload")
		}

		result, err := service.CreateSubmit(&input, userIdObjectId)

		if err != nil {
			return presenter.BadRequest(c, "Submission Failed")
		}

		return c.JSON(presenter.SubmissionSuccess(result))
	}
}
