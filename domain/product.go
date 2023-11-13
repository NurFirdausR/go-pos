package domain

type Product struct {
	Id          int
	Name        string
	PriceNet    int
	PriceGross  int
	StockQty    int
	Description string
	Image       string
	ExpDate     string
}
