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

	m := middleware.NewMiddleware(sessionUsecase)

	r.HandleFunc("/login", sessionDelivery.Login).Methods("POST")
	r.HandleFunc("/logout", sessionDelivery.Logout)
	r.HandleFunc("/signup", userDelivery.Create).Methods("POST")
	r.HandleFunc("/profile", userDelivery.GetUser).Methods("GET")
	r.HandleFunc("/profile", userDelivery.UpdateUser).Methods("PUT")
	r.HandleFunc("/profile/password", userDelivery.UpdatePassword).Methods("PUT")
	r.HandleFunc("/profile/avatar", userDelivery.UploadAvatar).Methods("POST")
	//r.HandleFunc("/profile", userDelivery.Delete).Methods("DELETE")
	r.HandleFunc("/pin", pinDelivery.Create).Methods("POST")
	r.HandleFunc("/pin/{id:[0-9]+}", pinDelivery.GetPin).Methods("GET")
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
