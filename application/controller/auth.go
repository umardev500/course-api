package controller

import (
	"course-api/application/service"
	"course-api/domain/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	validate *validator.Validate
	service  *service.AuthService
}

func NewAuthController(validate *validator.Validate, service *service.AuthService) *Auth {
	return &Auth{
		validate: validate,
		service:  service,
	}
}

// Login is handler for user login
//
// Return:
//   - error
func (a *Auth) Login(c *fiber.Ctx) error {
	var creds *model.LoginRequest = new(model.LoginRequest)
	if err := c.BodyParser(creds); err != nil {
		return failed(c, fiber.StatusBadRequest, err.Error())
	}

	// validation
	if err := a.validate.Struct(creds); err != nil {
		return failed(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	// call login service
	t, err := a.service.Login(*creds)
	if err != nil {
		if err == mongo.ErrNoDocuments || err == bcrypt.ErrMismatchedHashAndPassword {
			// check if error is no documents
			return failed(c, fiber.StatusNotFound, "user not founds")
		}

		return failed(c, fiber.StatusInternalServerError, err.Error())
	}
	return ok(c, fiber.StatusOK, "Login successfuly", model.LoginResponse{Token: *t})
}

// Register method is handler to handling user registration
//
// Return:
//   - error
func (a *Auth) Register(c *fiber.Ctx) error {
	var user *model.UserModel = new(model.UserModel)
	if err := c.BodyParser(user); err != nil {
		return failed(c, fiber.StatusBadRequest, err.Error())
	}

	// validate
	if err := a.validate.Struct(user); err != nil {
		return failed(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	if err := a.service.Register(*user); err != nil {
		return failed(c, fiber.StatusInternalServerError, err.Error())
	}

	return ok(c, fiber.StatusOK, "Register successfuly", nil)
}

func (a *Auth) Logout(c *fiber.Ctx) error {
	return nil
}
