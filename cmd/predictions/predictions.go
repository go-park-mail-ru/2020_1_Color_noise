package main

import (
	predictionsDeliveryGRPC "2020_1_Color_noise/internal/pkg/predictions/delivery/grpc"
	predictionsUsecase "2020_1_Color_noise/internal/pkg/predictions/usecase"


	"2020_1_Color_noise/internal/pkg/proto/predictions"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	us := predictionsUsecase.NewUsecase()

	predictionsDelivery := predictionsDeliveryGRPC.NewPredictionsService(us)

	lis, ok := net.Listen("tcp", ":8000")
	if ok != nil {
		log.Fatalln("cant listet port", ok)
	}

	server := grpc.NewServer()

	predictions.RegisterPredictionsServiceServer(server, predictionsDelivery)

	fmt.Println("starting server at :8000")
	server.Serve(lis)
}
