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

	if err != nil {
		panic("Some Error Occured")
	}

	app := fiber.New()

	// populate all routes
	routes.NewRoute(app, deps).SetupRoutes()

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
