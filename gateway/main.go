package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pnaskardev/pubjudge/gateway/config"
	"github.com/pnaskardev/pubjudge/gateway/routes"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("ENV variables not loaded")
	}
	deps, err := config.Init()

	app := fiber.New()

	// populate all routes
	routes.NewRoute(app, deps).SetupRoutes()

	log.Fatal(app.Listen(":3000"))
}
