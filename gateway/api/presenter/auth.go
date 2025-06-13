package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/pkg/entities"
)

func LoginSuccess(c *fiber.Ctx, user *entities.User, token string) error {
	user.Password = "" // Always strip password
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}

func LoginError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func BadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": msg,
	})
}
