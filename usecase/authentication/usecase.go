package authentication

import (
	"context"

	"github.com/NurFirdausR/go-pos/domain"
)

// UseCase ...
type UseCase interface {
	LoginHandler(ctx context.Context, auth domain.User) domain.User
	RegisterHandler(ctx context.Context, auth domain.User) domain.User
}
