package product

import (
	"context"

	"github.com/NurFirdausR/go-pos/domain"
	product_web "github.com/NurFirdausR/go-pos/web/product"
)

type UseCase interface {
	FindById(ctx context.Context, productId int) (domain.Product, error)
	FindAll(ctx context.Context) []domain.Product
	Save(ctx context.Context, request domain.Product) domain.Product
	Delete(ctx context.Context, request domain.Product)
	Update(ctx context.Context, request product_web.ProductUpdateRequest) domain.Product
}
