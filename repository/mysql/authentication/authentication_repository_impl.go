package authentication

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
)

type AuthenticateRepositoryImpl struct {
}

func NewAuthenticationRepository() AuthenticateRepository {
	return &AuthenticateRepositoryImpl{} // Return a pointer to the struct
}

func (r *AuthenticateRepositoryImpl) GetByName(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	query := "SELECT * from users where username = ?"
	rows, err := tx.QueryContext(ctx, query, username)
	helper.PanicIfError(err)

	defer rows.Close() // Close the rows when you're done with them

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil

	} else {
		return user, errors.New("user is not found")

	}

}

func (r *AuthenticateRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, auth domain.User) domain.User {
	query := "INSERT INTO users (username,password) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, query, auth.Username, auth.Password)
	helper.PanicIfError(err)
	userID, err := result.LastInsertId()
	if err != nil {
		helper.PanicIfError(err)
	}
	auth.Id = int(userID)
	return auth

}
