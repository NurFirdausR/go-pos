package category

type UpdateCategoryRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name" validate:"required,max=30"`
	Logo string `json:"logo" validate:"required"`
}
