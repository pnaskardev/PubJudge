package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pnaskardev/pubjudge/gateway/config"
	"github.com/pnaskardev/pubjudge/gateway/routes"
)

const idleTimeout = 5 * time.Second

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

	// log.Fatal(app.Listen(":3000"))
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	// This blocks the main thread until an interrupt is received
	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")

}
