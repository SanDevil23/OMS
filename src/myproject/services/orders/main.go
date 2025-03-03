package main

func main() {

	// http server
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()														// to run both the servers concurrently

	// grpc server
	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()
}