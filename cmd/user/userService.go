package main

import (
	userDeliveryGRPC "2020_1_Color_noise/internal/pkg/user/delivery/grpc"
	userRepository "2020_1_Color_noise/internal/pkg/user/repository"
	userUsecase "2020_1_Color_noise/internal/pkg/user/usecase"

	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	"2020_1_Color_noise/internal/pkg/proto/user"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.GetDBConfing()
	if err != nil {
		panic(err)
	}

	db := database.NewPgxDB()
	if err := db.Open(c); err != nil {
		panic(err)
	}

	userRepo := userRepository.NewRepo(db)
	userUse := userUsecase.NewUsecase(userRepo)
	userDelivery := userDeliveryGRPC.NewUserService(userUse)


	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	user.RegisterUserServiceServer(server, userDelivery)

	fmt.Println("starting server at :8000")
	server.Serve(lis)
}