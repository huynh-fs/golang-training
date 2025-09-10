package model

import "time"

type Order struct {
	ID              int64
	ProductID       int64
	QuantityOrdered int
	OrderStatus     string
	CreatedAt       time.Time
}	