package user_repository

import (
	"context"
	"crud-api/helper"
	"crud-api/model/domain"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepo UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `INSERT INTO "user" (name, email, password, username) VALUES ($1, $2, $3, $4) RETURNING id, name, email, username, password, created_at, updated_at`
	row := tx.QueryRowContext(ctx, SQL, user.Name, user.Email, user.Password, user.Username)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	helper.PanicIfError(err)

	return user
}

func (userRepo UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := `UPDATE "user" SET name=$1, email=$2, username=$3, password=$4 WHERE id=$5 RETURNING id, name, email, username, password, created_at, updated_at`
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Username, user.Password, user.Id)
	helper.PanicIfError(err)

	return user
}

func (userRepo UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := `DELETE FROM "user" WHERE id=$1`
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (userRepo UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := `SELECT id, name, email, username, password, created_at, updated_at FROM "user" WHERE id=$1`
	rows, err := tx.QueryContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
	defer rows.Close()

	user = domain.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (userRepo UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := `SELECT id, name, email, username, password, created_at, updated_at FROM "user"`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}
