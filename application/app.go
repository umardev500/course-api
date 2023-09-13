package application

import (
	"context"
	"course-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type App struct{}

func (a *App) Start(ctx context.Context) {
	app := fiber.New()
	port := os.Getenv("PORT")

	// load api routes
	routes.LoadAPIRoutes(app)

	// start the server
	log.Fatal(app.Listen(port))
}
