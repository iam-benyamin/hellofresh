package orderparam

type CreateOrderRequest struct {
	UserID      string `json:"user_id"`
	ProductCode string `json:"product_code"`
}

type SaveOrder struct {
	UserID           string
	ProductCode      string
	CustomerFullName string
	ProductName      string
	TotalAmount      float64
}
