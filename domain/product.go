package domain

type Product struct {
	Id          int    `json:"id" `
	Name        string `json:"name"`
	PriceNet    int    `json:"price_net"`
	PriceGross  int    `json:"price_gross"`
	StockQty    int    `json:"stock_qty"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ExpDate     string `json:"exp_date"`
	CompanyId   string `json:"company_id"`
	CategoryId  string `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
