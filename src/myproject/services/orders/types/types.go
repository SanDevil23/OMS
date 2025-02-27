package types

import (
	"Users/admin/Code/myProject/src/myproject/services/common/genproto/orders"
	"context"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error	
	GetOrders(context.Context) []*orders.Order
}