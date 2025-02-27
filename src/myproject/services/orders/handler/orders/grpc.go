package handler

import (
	"Users/admin/Code/myProject/src/myproject/services/common/genproto/orders"
	"Users/admin/Code/myProject/src/myproject/services/orders/types"
	"context"
)

type OrdersGrpcHandler struct {
	// service injection
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
	// unimplemented
}

func NewGrpcOrdersService() {
	gRPCHandler := &OrdersGrpcHandler{}
	// register the OrderServiceServer
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order {
		OrderId: 42,
		CustomerId: 2,
		ProuctId: 1,
		Quantity: 10,
	}

	err := h.ordersService.CreateOrder(ctx,  order)

	if err!=nil {
		return nil, err
	}	 

	res := &orders.CreateOrderResponse{
		Status : "success",
	}

	return res, nil
}