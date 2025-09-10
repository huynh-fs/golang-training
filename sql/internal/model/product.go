package model

import "time"

type Product struct {
	ID        int64
	Name      string
	Price     float64
	Quantity  int
	CreatedAt time.Time
}	