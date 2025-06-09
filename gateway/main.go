package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pnaskardev/pubjudge/gateway/routes"
)

func main() {


	

	app := fiber.New()

	// populate all routes
	routes.NewRoute(app).SetupRoutes()
	log.Fatal(app.Listen(":3000"))
}
