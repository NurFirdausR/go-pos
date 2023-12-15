package category

import (
	"context"
	"database/sql"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/repository/mysql/category"
	"github.com/go-playground/validator/v10"
)

type CategoryUsecase struct {
	CategoryRepository category.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryUsecase(categoryRepo category.CategoryRepository, DB *sql.DB, validate *validator.Validate) UseCase {
	return &CategoryUsecase{
		CategoryRepository: categoryRepo,
		DB:                 DB,
		Validate:           validate,
	}
}

func (usecase *CategoryUsecase) Save(ctx context.Context, request domain.Category) domain.Category {
	err := usecase.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	res := usecase.CategoryRepository.Save(ctx, tx, request)

	return res
}
