package authentication

import (
	"context"
	"database/sql"

	"github.com/NurFirdausR/go-pos/domain"
)

type AuthenticateRepository interface {
	GetByName(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	CreateUser(ctx context.Context, tx *sql.Tx, auth domain.User) domain.User
}
