package category

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/repository/mysql/category"
	category_web "github.com/NurFirdausR/go-pos/web/category"
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

func (usecase *CategoryUsecase) FindById(ctx context.Context, categoryId int) (domain.Category, error) {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categorys, err := usecase.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	return categorys, nil

}

func (usecase *CategoryUsecase) Update(ctx context.Context, request category_web.UpdateCategoryRequest) domain.Category {
	err := usecase.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	category, err := usecase.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)
	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02 15:04:05")

	category.Name = request.Name
	category.Logo = request.Logo
	category.UpdatedAt = dateString
	fmt.Println(category)
	res := usecase.CategoryRepository.Update(ctx, tx, category)
	return res
}
