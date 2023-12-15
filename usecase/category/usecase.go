package category

import (
	"context"

	"github.com/NurFirdausR/go-pos/domain"
)

type UseCase interface {
	Save(ctx context.Context, request domain.Category) domain.Category
}
