package domain

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,max=20"`
	Password string `json:"password" validate:"required,max=20"`
}
