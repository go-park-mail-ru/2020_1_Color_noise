package main

import (
	sessionDeliveryGRPC "2020_1_Color_noise/internal/pkg/session/delivery/grpc"
	sessionRepository "2020_1_Color_noise/internal/pkg/session/repository"
	sessionUsecase "2020_1_Color_noise/internal/pkg/session/usecase"

	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	"2020_1_Color_noise/internal/pkg/proto/session"
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

	sessionRepo := sessionRepository.NewRepo(db)
	sessionUse := sessionUsecase.NewUsecase(sessionRepo)
	sessionDelivery := sessionDeliveryGRPC.NewSessionManager(sessionUse)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	session.RegisterAuthSeviceServer(server, sessionDelivery)

	fmt.Println("starting server at :8000")
	server.Serve(lis)
}
