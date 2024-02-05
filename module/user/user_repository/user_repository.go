package user_repository

import (
	"context"
	"crud-api/model/domain"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
