package main

import (
	httpDeliverPin "pinterest/pkg/pin/delivery/http"
	repoPin "pinterest/pkg/pin/repository"
	usecasePin "pinterest/pkg/pin/usecase"

	httpDeliverSession "pinterest/pkg/session/delivery/http"
	repoSession "pinterest/pkg/session/repository"
	usecaseSession "pinterest/pkg/session/usecase"

	httpDeliverUser "pinterest/pkg/user/delivery/http"
	repoUser "pinterest/pkg/user/repository"
	usecaseUser "pinterest/pkg/user/usecase"

	middleware "pinterest/pkg/middleware"

	//"awesomeProject/internal/pkg/session/usecase"
	"github.com/gorilla/mux"
	"log"
	//"math/rand"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	userRepo := repoUser.NewUserRepo()
	userUsecase := usecaseUser.NewUserUsecase(userRepo)

	sessionRepo := repoSession.NewSessionRepo()
	sessionUsecase := usecaseSession.NewSessionUsecase(sessionRepo)
	sessionDelivery := httpDeliverSession.NewSessionHandler(sessionUsecase, userUsecase)

	userDelivery := httpDeliverUser.NewUserDelivery(userUsecase, sessionUsecase)

	pinRepo := repoPin.NewPinRepo()
	pinUsecase := usecasePin.NewPinUsecase(pinRepo)
	pinDelivery := httpDeliverPin.NewPinHandler(pinUsecase)

	m := middleware.NewMiddleware(sessionUsecase)

	r.HandleFunc("/login", sessionDelivery.Login).Methods("POST")
	r.HandleFunc("/logout", sessionDelivery.Logout)
	r.HandleFunc("/signup", userDelivery.AddUser).Methods("POST")
	r.HandleFunc("/profile/{id}", userDelivery.GetUser).Methods("GET")
	r.HandleFunc("/profile/{id}", userDelivery.UpdateUser).Methods("PUT")
	r.HandleFunc("/profile/{id}", userDelivery.DeleteUser).Methods("DELETE")
	r.HandleFunc("/pin", pinDelivery.Add).Methods("POST")
	r.HandleFunc("/pin/{id}", pinDelivery.GetPin).Methods("GET")
	//r.HandleFunc("/pin/{id}", pinDelivery.UpdatePin).Methods("PUT")
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
