package api

import (
	"github.com/johnsnoww/microservices/order/internal/application/core/domain"
	"github.com/johnsnoww/microservies/order/internal/ports"
)

type Application struct {
	db ports.DbPort
}

func NewApplication(db ports.DbPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreateOrder(order domain.Order) (domain.Order, error) {
	err := a.db.SaveOrder(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
