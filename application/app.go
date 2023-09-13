package application

import (
	"context"
	"course-api/config"
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
	router := routes.API{
		DB: config.NewMongo().Database("course"),
	}
	router.LoadAPIRoutes(app)

	// start the server
	log.Fatal(app.Listen(port))
}
