package main

import (
	predictionsDeliveryGRPC "2020_1_Color_noise/internal/pkg/predictions/delivery/grpc"


	"2020_1_Color_noise/internal/pkg/proto/predictions"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	predictionsDelivery := predictionsDeliveryGRPC.NewPredictionsService()

	lis, ok := net.Listen("tcp", ":8000")
	if ok != nil {
		log.Fatalln("cant listet port", ok)
	}

	server := grpc.NewServer()

	predictions.RegisterPredictionsServiceServer(server, predictionsDelivery)

	fmt.Println("starting server at :8000")
	server.Serve(lis)
}
