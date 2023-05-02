package products

import "time"

type Product struct {
	Title       string
	Description string
	CreatedAt   time.Time
	ID          string
	Price       int
}
