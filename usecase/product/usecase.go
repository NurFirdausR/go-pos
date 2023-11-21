package product

import (
	"context"
	"database/sql"

	"github.com/NurFirdausR/go-pos/domain"
)

type UseCase interface {
	Save(ctx context.Context, tx *sql.Tx, requset domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, requset domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, requset domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
