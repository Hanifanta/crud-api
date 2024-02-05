package app

import (
	"crud-api/module/user/user_controller"
	"crud-api/module/user/user_repository"
	"crud-api/module/user/user_service"
	"crud-api/utils"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	UserController user_controller.UserController
}

func NewController(config utils.Config) *Controller {
	db := NewDB(config.DBDriver, config.DBSource)
	validate := validator.New()

	userRepository := user_repository.NewUserRepository()

	userService := user_service.NewUserService(userRepository, db, validate)

	userController := user_controller.NewUserController(userService)

	return &Controller{UserController: userController}
}
