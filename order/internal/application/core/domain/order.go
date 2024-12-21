package domain

import (
	"time"
)

type OrderItem struct {
	ID        string
	ProductID string
	Quantity  int
	UnitPrice float32
}

type Order struct {
	ID         int64
	OrderItems []OrderItem
	Status     string
	CreatedAt  int64
	UpdatedAt  int64
}

func NewOrder(id int64, orderItems []OrderItem, status string) Order {
	return Order{
		ID:         id,
		OrderItems: orderItems,
		Status:     "PENDING",
		CreatedAt:  time.Now().Unix(),
		UpdatedAt:  time.Now().Unix(),
	}
}
