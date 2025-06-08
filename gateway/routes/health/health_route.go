package health

import "github.com/gofiber/fiber/v2"

type HealthCheckStruct struct {
	app *fiber.App
}

func NewHealthCheckRoutes(app *fiber.App) *HealthCheckStruct {
	return &HealthCheckStruct{app: app}
}

func (r *HealthCheckStruct) Register() {

	group := r.app.Group("/api")

	group.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
