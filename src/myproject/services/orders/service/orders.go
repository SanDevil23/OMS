package service

import (
	"Users/admin/Code/myProject/src/myproject/services/common/genproto/orders"
	"context"
)

// only business logic here
var ordersDb  = make([]*orders.Order, 0)


type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error{
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}