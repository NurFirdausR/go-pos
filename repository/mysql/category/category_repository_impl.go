package category

import (
	"context"
	"database/sql"

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
