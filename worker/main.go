package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pnaskardev/pubjudge/worker/config"
	"github.com/redis/go-redis/v9"
)

const (
	stream = "submission"
	group  = "submission_group"
)

var (
	consumer = uuid.New().String()
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

	// handle Stream Events
	// simple while loop
	for {
		messages, err := deps.Cache.Client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  []string{stream, ">"},
			Count:    50,
			// it will keep on listening to messages
			Block: 0,
			NoAck: false,
		}).Result()

		if err != nil {
			panic(err)
		}

		for _, stream := range messages {
			for _, msg := range stream.Messages {
				// Process the message
				fmt.Printf("Processing ID %s with values: %v\n", msg.ID, msg.Values)

				// Acknowledge after succesfull processing
				err := deps.Cache.Client.XAck(context.Background(), stream.Stream, group, msg.ID).Err()
				if err != nil {
					fmt.Printf("Failed to ack message %s: %v\n", msg.ID, err)
				}
			}
		}
	}

	// c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	// // This blocks the main thread until an interrupt is received
	// _ = <-c
	// fmt.Println("Gracefully shutting down...")
	// // _ = app.Shutdown()

	// fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// config.CloseDBConnection()
	// config.CloseCacheConnection()
}
