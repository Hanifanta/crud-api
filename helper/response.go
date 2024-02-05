package helper

import (
	"crud-api/model/domain"
	"crud-api/model/web/user"
)

func ToUserResponse(userData domain.User) user.UserResponse {
	return user.UserResponse{
		Id:        userData.Id,
		Email:     userData.Email,
		Username:  userData.Username,
		Name:      userData.Name,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
}

func ToUserResponses(usersData []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, userData := range usersData {
		userResponses = append(userResponses, ToUserResponse(userData))
	}
	return userResponses
}
