package product

import (
	"context"

	"github.com/NurFirdausR/go-pos/domain"
)

type UseCase interface {
	FindById(ctx context.Context, productId int) (domain.Product, error)
	FindAll(ctx context.Context) []domain.Product
	Save(ctx context.Context, requset domain.Product) domain.Product
	Delete(ctx context.Context, requset domain.Product)
	Update(ctx context.Context, requset domain.Product) domain.Product
}
