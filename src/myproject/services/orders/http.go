package main

import (
	handler "Users/admin/Code/myProject/src/myproject/services/orders/handler/orders"
	"Users/admin/Code/myProject/src/myproject/services/orders/service"
	"log"
	"net/http"
)

type httpServer struct {
	// port number
	addr string
}

func NewHttpServer(addr string) *httpServer {
	// creating a new http server on the port -- addr
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)

	// registering the server mux -- router
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)
	

	/*
		here we have started the http serve om the [port] 
		and passed http.handler as the server mux --- router
	*/
	return http.ListenAndServe(s.addr, router)
}