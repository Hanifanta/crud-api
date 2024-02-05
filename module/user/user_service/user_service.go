package user_service

import (
	"context"
	"crud-api/model/web/user"
)

type UserService interface {
	Create(ctx context.Context, userRequest user.UserRequest) user.UserResponse
	Update(ctx context.Context, Id string, userRequest user.UserRequest) user.UserResponse
	Delete(ctx context.Context, Id string)
	FindById(ctx context.Context, Id string) user.UserResponse
	FindAll(ctx context.Context) []user.UserResponse
}
