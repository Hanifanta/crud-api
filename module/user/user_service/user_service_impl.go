package user_service

import (
	"context"
	"crud-api/helper"
	"crud-api/model/domain"
	"crud-api/model/web/user"
	"crud-api/module/user/user_repository"
	"crud-api/utils"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository user_repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, DB: DB, Validate: validate}
}

func (userService UserServiceImpl) Create(ctx context.Context, userRequest user.UserRequest) user.UserResponse {
	err := userService.Validate.Struct(userRequest)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	hashedPassword, err := utils.HashPassword(userRequest.Password)
	helper.PanicIfError(err)

	requestDomain := domain.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Username: userRequest.Username,
		Password: hashedPassword,
	}

	user := userService.UserRepository.Create(ctx, tx, requestDomain)

	return helper.ToUserResponse(user)
}

func (userService UserServiceImpl) Update(ctx context.Context, Id string, userRequest user.UserRequest) user.UserResponse {
	err := userService.Validate.Struct(userRequest)
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindById(ctx, tx, domain.User{Id: Id})
	helper.PanicIfError(err)

	requestDomain := domain.User{
		Id:       user.Id,
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Username: userRequest.Username,
	}

	update := userService.UserRepository.Update(ctx, tx, requestDomain)

	return helper.ToUserResponse(update)
}

func (userService UserServiceImpl) Delete(ctx context.Context, Id string) {
	err := userService.Validate.Var(Id, "required")
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindById(ctx, tx, domain.User{Id: Id})
	helper.PanicIfError(err)

	userService.UserRepository.Delete(ctx, tx, user)
}

func (userService UserServiceImpl) FindById(ctx context.Context, Id string) user.UserResponse {
	err := userService.Validate.Var(Id, "required")
	helper.PanicIfError(err)

	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := userService.UserRepository.FindById(ctx, tx, domain.User{Id: Id})
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (userService UserServiceImpl) FindAll(ctx context.Context) []user.UserResponse {
	tx, err := userService.DB.Begin()
	helper.PanicIfError(err)

	users := userService.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}
