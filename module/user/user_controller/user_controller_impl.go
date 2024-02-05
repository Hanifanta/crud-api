package user_controller

import (
	"crud-api/helper"
	"crud-api/model/web"
	"crud-api/model/web/user"
	"crud-api/module/user/user_service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService user_service.UserService
}

func NewUserController(userService user_service.UserService) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (userController UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := user.UserRequest{}

	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := userController.UserService.Create(request.Context(), userCreateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := user.UserRequest{}

	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userResponse := userController.UserService.Update(request.Context(), params.ByName("id"), userUpdateRequest)

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userController.UserService.Delete(request.Context(), params.ByName("id"))

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := userController.UserService.FindById(request.Context(), params.ByName("id"))

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}

func (userController UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := userController.UserService.FindAll(request.Context())

	baseResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, baseResponse)
}
