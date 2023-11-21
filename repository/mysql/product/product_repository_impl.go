package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// Name        string
//
//	PriceNet    int
//	PriceGross  int
//	StockQty    int
//	Description string
//	Image       string
//	ExpDate     string
func (r *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(name,price_net,price_gross,stock_qty,description,image,exp_date) VALUES(?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.PriceNet, product.PriceGross, product.StockQty, product.Description, product.Image, product.ExpDate)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	product.Id = int(id)
	return product

}
func (r *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name=?,price_net=?,price_gross=?,stock_qty=?,description=?,image=?,exp_date=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.PriceNet, product.PriceGross, product.StockQty, product.Description, product.Image, product.ExpDate, product.Id)
	helper.PanicIfError(err)
	return product
}
func (r *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT * FROM products WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()
	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.PriceNet, &product.PriceGross, &product.StockQty, &product.Description, &product.Image, &product.ExpDate)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("Product is not found")
	}

}
func (r *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "DELETE * FROM products where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}
func (r *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT * FROM products"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.PriceNet, &product.PriceGross, &product.StockQty, &product.Description, &product.Image, &product.ExpDate)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	err = rows.Err()
	helper.PanicIfError(err)
	return products

}
