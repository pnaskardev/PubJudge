package router_types

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/config"
)

type Router struct {
	App  *fiber.App
	Api  fiber.Router
	Deps *config.App
}
