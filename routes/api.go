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

// LoadAPIRoutes load all api routes
func (api *API) LoadAPIRoutes(app *fiber.App) {
	router := app.Group("/api")
	router.Use(middleware.NewLogger())
	router.Route("/auth", api.loadAuthRoutes)
	router.Route("/video", api.loadVideoRoutes)
	router.Route("/upload", api.loadUploadRoutes)
}

// loadAuthRoutes load all routes of auth
func (api *API) loadAuthRoutes(router fiber.Router) {
	collection := api.DB.Collection("users")
	authRepo := repository.NewAuthRepository(collection)
	authService := service.NewAuthService(authRepo)
	handler := controller.NewAuthController(config.Validate, authService)

	router.Post("/login", middleware.NewLimiter(), handler.Login)
	router.Post("/register", handler.Register)
	router.Post("/logout", handler.Logout)
}

// loadUploadRoutes - load all routes of upload
//
// Params:
//   - router fiber.Router
//
// Return:
//   - void
func (api *API) loadUploadRoutes(router fiber.Router) {
	uploadService := service.NewUploadService()
	handler := controller.NewUploadController(uploadService)
	router.Post("/chunk/:file_id/:total/:index", handler.UploadChunk)
}

// loadVideoRoutes - load all routes of video handler
//
// Params:
//   - router fiber.Router
//
// Return:
//   - void
func (api *API) loadVideoRoutes(router fiber.Router) {
	handler := controller.NewVideoController()
	router.Post("/upload", handler.Upload)
}
