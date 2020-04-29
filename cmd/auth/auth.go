package main

import (
	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	"2020_1_Color_noise/internal/pkg/session"
	"2020_1_Color_noise/internal/pkg/session/delivery/grpc"
	"2020_1_Color_noise/internal/pkg/session/repository/repo"
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

	c.User = "postgres"
	c.Password = "password"

	db := database.NewPgxDB()
	if err := db.Open(c); err != nil {
		panic(err)
	}

	sessionRepo := repo.NewRepo(db)

	lis, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	session.RegisterAuthCheckerServer(server, rpc.NewSessionManager(sessionRepo))

	fmt.Println("starting server at :8003")
	server.Serve(lis)
}
