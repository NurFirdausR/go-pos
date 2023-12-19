package product

import (
	"context"
	"database/sql"
	"time"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
	"github.com/NurFirdausR/go-pos/repository/mysql/product"
	product_web "github.com/NurFirdausR/go-pos/web/product"
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

func (usecase *ProductUseCase) Update(ctx context.Context, request product_web.ProductUpdateRequest) domain.Product {
	err := usecase.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	product, err := usecase.ProductRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)
	currentTime := time.Now()
	dateString := currentTime.Format("2006-01-02 15:04:05")

	product.Name = request.Name
	product.PriceNet = request.PriceNet
	product.PriceGross = request.PriceGross
	product.StockQty = request.StockQty
	product.Description = request.Description
	product.Image = request.Image
	product.ExpDate = request.ExpDate
	product.UpdatedAt = dateString
	product.CategoryId = request.CategoryId
	product.CompanyId = request.CompanyId
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
