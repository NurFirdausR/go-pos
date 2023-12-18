package category

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NurFirdausR/go-pos/domain"
	"github.com/NurFirdausR/go-pos/helper"
)

type CategoryRepostoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepostoryImpl{}
}

func (r *CategoryRepostoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name,logo) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.Logo)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (r *CategoryRepostoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT * FROM categories WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.Logo, &category.CreatedAt, &category.UpdatedAt)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}
