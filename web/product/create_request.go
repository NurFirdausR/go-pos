package product

import "time"

type ProductCreateRequest struct {
	Id          int       `json:"id" `
	Name        string    `json:"name" validate:"required"`
	PriceNet    int       `json:"price_net" validate:"required"`
	PriceGross  int       `json:"price_gross" validate:"required"`
	StockQty    int       `json:"stock_qty" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	ExpDate     string    `json:"exp_date" validate:"required"`
	CompanyId   string    `json:"company_id" validate:"required"`
	CategoryId  string    `json:"category_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
