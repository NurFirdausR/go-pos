package category

import (
	"context"

	"github.com/NurFirdausR/go-pos/domain"
	category_web "github.com/NurFirdausR/go-pos/web/category"
)

type UseCase interface {
	Save(ctx context.Context, request domain.Category) domain.Category
	FindById(ctx context.Context, categoryId int) (domain.Category, error)
	Update(ctx context.Context, request category_web.UpdateCategoryRequest) domain.Category
}
