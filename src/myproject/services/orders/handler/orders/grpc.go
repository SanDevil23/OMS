package handler

import (
	"Users/admin/Code/myProject/src/myproject/services/common/genproto/orders"
	"Users/admin/Code/myProject/src/myproject/services/orders/types"
	"context"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// service injection
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
	// unimplemented
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}
	// register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler);
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	ord := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrdersResponse{
		Orders:  ord,
	}
	return res, nil
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order {
		OrderId: 42,
		CustomerId: 2,
		ProductId: 1,
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