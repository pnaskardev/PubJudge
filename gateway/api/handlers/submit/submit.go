package submission_handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/api/presenter"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
	"github.com/pnaskardev/pubjudge/gateway/pkg/submit"
)

func HandleSubmit(service submit.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input entities.Submission

		if err := c.BodyParser(&input); err != nil {
			return presenter.BadRequest(c, "Inavlid Payload")
		}

		submission_instance, err := service.CreateSubmit(&input)

		if err != nil {
			return presenter.BadRequest(c, "Submission Failed")
		}

		return c.JSON(presenter.SubmissionSuccess(submission_instance))
	}
}
