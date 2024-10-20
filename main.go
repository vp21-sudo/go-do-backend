package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vp21-sudo/go-do-backend/db"
	"github.com/vp21-sudo/go-do-backend/routes"
)

func main() {
	app := fiber.New()

	db.Connect()

	routes.RegisterRoutes(app)

	go func() {
		if err := app.Listen(":9000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.Disconnect(ctx)

	log.Println("Fiber was successfully shutdown.")
}
