package service

import (
	"Users/admin/Code/myProject/src/myproject/services/common/genproto/orders"
	"context"
)

// only business logic here
var orders1  = make([]*orders.Order, 0)
type OrderService struct {
	// store
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order){
	orders1 = append(orders1, order)
	return nil
}