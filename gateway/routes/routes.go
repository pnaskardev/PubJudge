package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/routes/health"
	"github.com/pnaskardev/pubjudge/gateway/routes/submit"
)

type Router struct {
	app *fiber.App
}

func NewRoute(App *fiber.App) *Router {
	return &Router{app: App}
}

func (r *Router) SetupRoutes() {
	health.NewHealthCheckRoutes(r.app).Register()
	submit.NewSubmissionRoutes(r.app).Register()
}
