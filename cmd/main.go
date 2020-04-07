package main

import (
	boardDeliveryHttp "2020_1_Color_noise/internal/pkg/board/delivery/http"
	boardRepo "2020_1_Color_noise/internal/pkg/board/repository"
	boardUsecase "2020_1_Color_noise/internal/pkg/board/usecase"
	notificationsDeliveryHttp "2020_1_Color_noise/internal/pkg/notifications/delivery/http"
	notificationsRepo "2020_1_Color_noise/internal/pkg/notifications/repository"
	notificationsUsecase "2020_1_Color_noise/internal/pkg/notifications/usecase"

	commentDeliveryHttp "2020_1_Color_noise/internal/pkg/comment/delivery/http"
	commentRepo "2020_1_Color_noise/internal/pkg/comment/repository"
	commentUsecase "2020_1_Color_noise/internal/pkg/comment/usecase"

	"2020_1_Color_noise/internal/pkg/database"

	pinDeliveryHttp "2020_1_Color_noise/internal/pkg/pin/delivery/http"
	pinRepo "2020_1_Color_noise/internal/pkg/pin/repository"
	pinUsecase "2020_1_Color_noise/internal/pkg/pin/usecase"

	sessionDeliveryHttp "2020_1_Color_noise/internal/pkg/session/delivery/http"
	sessionRepo "2020_1_Color_noise/internal/pkg/session/repository"
	sessionUsecase "2020_1_Color_noise/internal/pkg/session/usecase"

	userDeliveryHttp "2020_1_Color_noise/internal/pkg/user/delivery/http"
	userRepo "2020_1_Color_noise/internal/pkg/user/repository"
	userUsecase "2020_1_Color_noise/internal/pkg/user/usecase"

	listDeliveryHttp "2020_1_Color_noise/internal/pkg/list/delivery/http"
	listRepo "2020_1_Color_noise/internal/pkg/list/repository"
	listUsecase "2020_1_Color_noise/internal/pkg/list/usecase"

	searchHandler "2020_1_Color_noise/internal/pkg/search"

	"2020_1_Color_noise/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	db := database.NewPgxDB()
	if err := db.Open(); err != nil {
		panic(err)
	}

	userRepo := userRepo.NewRepo(db)
	userUsecase := userUsecase.NewUsecase(userRepo)

	sessionRepo := sessionRepo.NewRepo(db)
	sessionUsecase := sessionUsecase.NewUsecase(sessionRepo)
	sessionDelivery := sessionDeliveryHttp.NewHandler(sessionUsecase, userUsecase)

	userDelivery := userDeliveryHttp.NewHandler(userUsecase, sessionUsecase)

	pinRepo := pinRepo.NewRepo(db)
	pinUsecase := pinUsecase.NewUsecase(pinRepo)
	pinDelivery := pinDeliveryHttp.NewHandler(pinUsecase)

	boardRepo := boardRepo.NewRepo(db)
	boardUsecase := boardUsecase.NewUsecase(boardRepo)
	boardDelivery := boardDeliveryHttp.NewHandler(boardUsecase)

	commentRepo := commentRepo.NewRepo(db)
	commentUsecase := commentUsecase.NewUsecase(commentRepo)
	commentDelivery := commentDeliveryHttp.NewHandler(commentUsecase)

	listRepo := listRepo.NewRepo(db)
	listUsecase := listUsecase.NewUsecase(listRepo)
	listDelivery := listDeliveryHttp.NewHandler(listUsecase)

	notificationsRepo := notificationsRepo.NewRepo(db)
	notificationsUsecase := notificationsUsecase.NewUsecase(notificationsRepo)
	notificationsDelivery := notificationsDeliveryHttp.NewHandler(notificationsUsecase)

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
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.DeletePin).Methods("DELETE")

	r.HandleFunc("/api/board", boardDelivery.Create).Methods("POST")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.Update).Methods("PUT")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.Delete).Methods("DELETE")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.GetBoard).Methods("GET")
	r.HandleFunc("/api/board/name/user/{id:[0-9]+}", boardDelivery.GetBoard).Methods("GET")
	r.HandleFunc("/api/board/user/{id:[0-9]+}", boardDelivery.Fetch).Methods("GET")

	r.HandleFunc("/api/comment", commentDelivery.Create).Methods("POST")
	r.HandleFunc("/api/comment/{id:[0-9]+}", commentDelivery.GetComment).Methods("GET")
	r.HandleFunc("/api/comment/pin/{id:[0-9]+}", commentDelivery.Fetch).Methods("GET")

	r.HandleFunc("/api/search", searchHandler.Search).Methods("GET")

	r.HandleFunc("/api/list", listDelivery.GetMainList).Methods("GET")
	r.HandleFunc("/api/list/sub", listDelivery.GetSubList).Methods("GET")
	r.HandleFunc("/api/list/recommendation", listDelivery.GetRecommendationList).Methods("GET")

	r.HandleFunc("/api/notifications", notificationsDelivery.GetNotifications).Methods("GET")

	r.Use(m.CORSMiddleware)
	r.Use(m.AuthMiddleware)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
