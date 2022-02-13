package main

import (
	"grpc-unary-medium/payment/paymentpb"
	"grpc-unary-medium/payment/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// create service 
	paymentService := service.NewPaymentService()

	// create grpc server
	grpcServer := grpc.NewServer()

	// mendaftarkan service
	paymentpb.RegisterPaymentServer(grpcServer, paymentService)

	// membuat listener pada tcp di port 3000
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starting...")
	// start server
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
