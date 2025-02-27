package handler

import (
	"Users/admin/Code/myProject/src/myproject/services/common/genproto/orders"
	"Users/admin/Code/myProject/src/myproject/services/common/util"
	"Users/admin/Code/myProject/src/myproject/services/orders/types"
	"net/http"
	// "github.com/sikozonpc/kitchen/services/common/genproto/orders"
	// "github.com/sikozonpc/kitchen/services/common/util"
	// "github.com/sikozonpc/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}


// in-point for the orders
func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderId: 42,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQty(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}