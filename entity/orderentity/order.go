package orderentity

import "time"

type Order struct {
	ID               string
	UserID           string
	ProductCode      string
	CustomerFullName string
	ProductName      string
	TotalAmount      float64
	CreatedAt        time.Time
}
