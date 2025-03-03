package main

import (
	handler "Users/admin/Code/myProject/src/myproject/services/orders/handler/orders"
	"Users/admin/Code/myProject/src/myproject/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

// this a constructor
func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}


/* 	s is a gRPCServer object	
	which has the addr of the port '
	on which the server will be running
*/
func (s *gRPCServer) Run() error {
	
	// creating a tcp connection
	lis, err := net.Listen("tcp", s.addr);

	// handling any error found while establishing tcp connection
	if err!=nil{
		log.Fatalf("failed to listen : %v", err)
	}
	
	// initializing a server
	gRPCServer := grpc.NewServer()


	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService); 
	log.Println("Starting grpc server on ", s.addr)
	return gRPCServer.Serve(lis)

}