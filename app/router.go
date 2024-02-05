package app

import (
	"crud-api/utils"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(config utils.Config) *httprouter.Router {
	controller := NewController(config)

	router := httprouter.New()

	router.POST("/user", controller.UserController.Create)
	router.PUT("/user/:id", controller.UserController.Update)
	router.DELETE("/user/:id", controller.UserController.Delete)
	router.GET("/user/:id", controller.UserController.FindById)
	router.GET("/user", controller.UserController.FindAll)

	return router
}
