package main

import (
	//"echo_example/user/usecase"
	config "2020_1_Color_noise/internal/pkg/config"
	pinDeliveryHttp "2020_1_Color_noise/internal/pkg/pin/delivery/http"
	pinRepo "2020_1_Color_noise/internal/pkg/pin/repository"
	pinUsecase "2020_1_Color_noise/internal/pkg/pin/usecase"

	sessionDeliveryHttp "2020_1_Color_noise/internal/pkg/session/delivery/http"
	sessionRepo "2020_1_Color_noise/internal/pkg/session/repository"
	sessionUsecase "2020_1_Color_noise/internal/pkg/session/usecase"

	userDeliveryHttp "2020_1_Color_noise/internal/pkg/user/delivery/http"
	userRepo "2020_1_Color_noise/internal/pkg/user/repository"
	userUsecase "2020_1_Color_noise/internal/pkg/user/usecase"

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

	config.Start()

	userRepo := userRepo.NewRepo()
	userUsecase := userUsecase.NewUsecase(userRepo)

	sessionRepo := sessionRepo.NewRepo()
	sessionUsecase := sessionUsecase.NewUsecase(sessionRepo)
	sessionDelivery := sessionDeliveryHttp.NewHandler(sessionUsecase, userUsecase)

	userDelivery := userDeliveryHttp.NewHandler(userUsecase, sessionUsecase)

	pinRepo := pinRepo.NewRepo()
	pinUsecase := pinUsecase.NewUsecase(pinRepo)
	pinDelivery := pinDeliveryHttp.NewHandler(pinUsecase)

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
	//r.HandleFunc("/api/user", userDelivery.Delete).Methods("DELETE")
	r.HandleFunc("/api/pin", pinDelivery.Create).Methods("POST")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.GetPin).Methods("GET")
	r.HandleFunc("/api/pin/user/{id:[0-9]+}", pinDelivery.Fetch).Methods("GET")
	//r.HandleFunc("/pin/{id}", pinDelivery.Update).Methods("PUT")
	//r.HandleFunc("/pin/{id}", pinDelivery.DeletePin).Methods("DELETE")
	r.Use(m.AuthMiddleware)
	r.Use(m.CORSMiddleware)
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
