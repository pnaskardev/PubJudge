package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/pnaskardev/pubjudge/worker/config"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("ENV variables not loaded")
	}
	// deps, err := config.Init()
	// if err != nil {
	// 	panic("Some Error Occured")
	// }

	// handle Stream Events

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	// This blocks the main thread until an interrupt is received
	_ = <-c
	fmt.Println("Gracefully shutting down...")
	// _ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// config.CloseDBConnection()
	config.CloseCacheConnection()
	fmt.Println("Fiber was successful shutdown.")

	fmt.Println("Hello World")
}
