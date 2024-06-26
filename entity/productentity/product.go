package productentity

import "time"

type Product struct {
	ID        string
	Name      string
	Price     float64
	Code      string
	CreatedAt time.Time
}
