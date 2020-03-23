package main

import (
	//"echo_example/user/usecase"
	pinDeliveryHttp "2020_1_Color_noise/internal/pkg/pin/delivery/http"
	pinRepo "2020_1_Color_noise/internal/pkg/pin/repository"
	pinUsecase "2020_1_Color_noise/internal/pkg/pin/usecase"

	sessionDeliveryHttp "2020_1_Color_noise/internal/pkg/session/delivery/http"
	sessionRepo "2020_1_Color_noise/internal/pkg/session/repository"
	sessionUsecase "2020_1_Color_noise/internal/pkg/session/usecase"

	userDeliveryHttp "2020_1_Color_noise/internal/pkg/user/delivery/http"
	userRepo "2020_1_Color_noise/internal/pkg/user/repository"
	userUsecase "2020_1_Color_noise/internal/pkg/user/usecase"

	boardDeliveryHttp "2020_1_Color_noise/internal/pkg/board/delivery/http"
	boardRepo "2020_1_Color_noise/internal/pkg/board/repository"
	boardUsecase "2020_1_Color_noise/internal/pkg/board/usecase"

	commentDeliveryHttp "2020_1_Color_noise/internal/pkg/comment/delivery/http"
	commentRepo "2020_1_Color_noise/internal/pkg/comment/repository"
	commentUsecase "2020_1_Color_noise/internal/pkg/comment/usecase"

	searchHandler "2020_1_Color_noise/internal/pkg/search"

	"2020_1_Color_noise/internal/pkg/middleware"

	//"awesomeProject/internal/pkg/session/usecase"
	"github.com/gorilla/mux"
	"log"
	//"math/rand"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	userRepo := userRepo.NewRepo()
	userUsecase := userUsecase.NewUsecase(userRepo)

	sessionRepo := sessionRepo.NewRepo()
	sessionUsecase := sessionUsecase.NewUsecase(sessionRepo)
	sessionDelivery := sessionDeliveryHttp.NewHandler(sessionUsecase, userUsecase)

	userDelivery := userDeliveryHttp.NewHandler(userUsecase, sessionUsecase)

	pinRepo := pinRepo.NewRepo()
	pinUsecase := pinUsecase.NewUsecase(pinRepo)
	pinDelivery := pinDeliveryHttp.NewHandler(pinUsecase)

	boardRepo := boardRepo.NewRepo()
	boardUsecase := boardUsecase.NewUsecase(boardRepo)
	boardDelivery := boardDeliveryHttp.NewHandler(boardUsecase)

	commentRepo := commentRepo.NewRepo()
	commentUsecase := commentUsecase.NewUsecase(commentRepo)
	commentDelivery := commentDeliveryHttp.NewHandler(commentUsecase)

	searchHandler := searchHandler.NewHandler(commentUsecase, pinUsecase, userUsecase)

	m := middleware.NewMiddleware(sessionUsecase)

	r.HandleFunc("/api/auth", sessionDelivery.Login).Methods("POST")
	r.HandleFunc("/api/auth", sessionDelivery.Logout).Methods("DELETE")
	r.HandleFunc("/api/user", userDelivery.Create).Methods("POST")
	r.HandleFunc("/api/user", userDelivery.GetUser).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}", userDelivery.GetOtherUser).Methods("GET")
	r.HandleFunc("/api/user/settings/profile", userDelivery.UpdateProfile).Methods("PUT")
	r.HandleFunc("/api/user/settings/description", userDelivery.UpdateDescription).Methods("PUT")
	r.HandleFunc("/api/user/settings/password", userDelivery.UpdatePassword).Methods("PUT")
	r.HandleFunc("/api/user/settings/avatar", userDelivery.UploadAvatar).Methods("PUT")
	r.HandleFunc("/api/user/following/{id:[0-9]+}", userDelivery.Follow).Methods("POST")
	r.HandleFunc("/api/user/following/{id:[0-9]+}", userDelivery.Unfollow).Methods("DELETE")
	r.HandleFunc("/api/user/subscriptions/{id:[0-9]+}", userDelivery.GetSubscribtions).Methods("GET")
	r.HandleFunc("/api/user/subscribers/{id:[0-9]+}", userDelivery.GetSubscribers).Methods("GET")
	//r.HandleFunc("/api/user", userDelivery.Delete).Methods("DELETE")
	r.HandleFunc("/api/pin", pinDelivery.Create).Methods("POST")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.GetPin).Methods("GET")
	r.HandleFunc("/api/pin/user/{id:[0-9]+}", pinDelivery.Fetch).Methods("GET")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.Update).Methods("PUT")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.Update).Methods("DELETE")
	r.HandleFunc("/api/board", boardDelivery.Create).Methods("POST")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.GetBoard).Methods("GET")
	r.HandleFunc("/api/board/user/{id:[0-9]+}", boardDelivery.Fetch).Methods("GET")
	r.HandleFunc("/api/comment", commentDelivery.Create).Methods("POST")
	r.HandleFunc("/api/comment/{id:[0-9]+}", commentDelivery.GetComment).Methods("GET")
	r.HandleFunc("/api/comment/pin/{id:[0-9]+}", commentDelivery.Fetch).Methods("GET")
	r.HandleFunc("/api/search", searchHandler.Search).Methods("GET")

	r.Use(m.CORSMiddleware)
	r.Use(m.AuthMiddleware)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
