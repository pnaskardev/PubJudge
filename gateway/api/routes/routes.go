package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pnaskardev/pubjudge/gateway/api/routes/auth"
	"github.com/pnaskardev/pubjudge/gateway/api/routes/health"
	"github.com/pnaskardev/pubjudge/gateway/api/routes/submit"
	"github.com/pnaskardev/pubjudge/gateway/api/routes/user"
	"github.com/pnaskardev/pubjudge/gateway/config"
	"github.com/pnaskardev/pubjudge/gateway/types/router_types"
)

type Router struct {
	Router *router_types.Router
}

func NewRoute(App *fiber.App, Deps *config.App) *Router {

	api := App.Group("/api", logger.New())

	defaultRouter := router_types.Router{App: App, Deps: Deps, Api: api}

	return &Router{Router: &defaultRouter}
}

func (r *Router) SetupRoutes() {
	// Health check does not needs any external dependencies

	health.NewHealthCheckRoutes(r.Router).Register()

	auth.NewAuthRoutes(r.Router).Register()

	user.NewUserRoutes(r.Router).Register()

	submit.NewSubmitRoutes(r.Router).Register()
}
