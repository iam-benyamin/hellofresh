package productparam

type ProductRequest struct {
	ProductCode string
}

type ProductResponse struct {
	ID    string
	Name  string
	Code  string
	Price float64
}
