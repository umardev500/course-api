package routes

import (
	"course-api/application/controller"
	"course-api/application/repository"
	"course-api/application/service"
	"course-api/config"
	"course-api/middleware"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type API struct {
	DB *mongo.Database
}

func (api *API) LoadAPIRoutes(app *fiber.App) {
	router := app.Group("/api")
	router.Use(middleware.NewLogger())
	router.Route("/auth", api.loadAuthRoutes)
}

func (api *API) loadAuthRoutes(router fiber.Router) {
	collection := api.DB.Collection("users")
	authRepo := repository.NewAuthRepository(collection)
	authService := service.NewAuthService(authRepo)
	handler := controller.NewAuthController(config.Validate, authService)

	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
	router.Post("/logout", handler.Logout)
}
