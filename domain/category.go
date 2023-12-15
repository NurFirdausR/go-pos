package domain

type Category struct {
	Id        int    `json:"id"`
	Name      string `json:"name" validate:"required,max=30"`
	Logo      string `json:"logo" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
