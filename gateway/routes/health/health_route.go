package health

import (
	"github.com/gofiber/fiber/v2"

	router_types "github.com/pnaskardev/pubjudge/gateway/types/router_types"
)

type HealthRouter struct {
	Router *router_types.Router
}

func NewHealthCheckRoutes(router *router_types.Router) *HealthRouter {
	return &HealthRouter{Router: router}
}

func (r *HealthRouter) Register() {

	r.Router.Api.Get("/ping", func(c *fiber.Ctx) error {

		// r.Router.Deps.Db.Database.Collection("test").Aggregate([$match:{}])

		return c.SendStatus(fiber.StatusOK)
	})
}
