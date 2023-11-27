package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/repository/mysql/product"
	"github.com/go-playground/validator/v10"
)

type ProductUseCase struct {
	ProductRepository product.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductUsecase(productRepo product.ProductRepository, DB *sql.DB, validate *validator.Validate) UseCase {
	return &ProductUseCase{
		ProductRepository: productRepo,
		DB:                DB,
		Validate:          validate,
	}
}

func (usecase *ProductUseCase) Save(ctx context.Context, request domain.Product) domain.Product {
	err := usecase.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	res := usecase.ProductRepository.Save(ctx, tx, request)

	return res

}

func (usecase *ProductUseCase) Update(ctx context.Context, request domain.Product) domain.Product {
	err := usecase.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product, err := usecase.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		helper.PanicIfError(err)
	}
	fmt.Println(product)
	res := usecase.ProductRepository.Update(ctx, tx, product)
	return res
}

func (usecase *ProductUseCase) Delete(ctx context.Context, request domain.Product) {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := usecase.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		helper.PanicIfError(err)
	}
	usecase.ProductRepository.Delete(ctx, tx, product)
}

func (usecase *ProductUseCase) FindById(ctx context.Context, productId int) (domain.Product, error) {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := usecase.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)
	return product, nil
}

func (usecase *ProductUseCase) FindAll(ctx context.Context) []domain.Product {
	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	products := usecase.ProductRepository.FindAll(ctx, tx)
	var ProductsData []domain.Product
	for _, prod := range products {
		ProductsData = append(ProductsData, domain.Product(prod))
	}
	return ProductsData
}
