package domain

type Product struct {
	Id          int    `json:"id" `
	Name        string `json:"name" validate:"required"`
	PriceNet    int    `json:"price_net" validate:"required"`
	PriceGross  int    `json:"price_gross" validate:"required"`
	StockQty    int    `json:"stock_qty" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image" validate:"required"`
	ExpDate     string `json:"exp_date" validate:"required"`
	CompanyId   int    `json:"company_id" validate:"required"`
	CategoryId  int    `json:"category_id" validate:"required"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
