package domain

type OrderType string

const (
	Buy  OrderType = "BUY"
	Sell OrderType = "SELL"
)

type Order struct {
	ID       	int
	Type     	OrderType
	Amount		int
	Price  		float64
}