package main

import (
	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	"2020_1_Color_noise/internal/pkg/middleware"
	sessionRepository "2020_1_Color_noise/internal/pkg/session/repository"
	sessionUsecase "2020_1_Color_noise/internal/pkg/session/usecase"
	"log"

	chatDeliveryHttp "2020_1_Color_noise/internal/pkg/chat/delivery/http"
	chatRepository "2020_1_Color_noise/internal/pkg/chat/repository"
	chatUsecase "2020_1_Color_noise/internal/pkg/chat/usecase"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"

	"2020_1_Color_noise/internal/pkg/chat/delivery"
)

func main() {

	r := mux.NewRouter()

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

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	hub := delivery.NewHub()
	go hub.Run()

	chatRepo := chatRepository.NewRepository(db)
	chatUse := chatUsecase.NewUsecase(chatRepo)
	chatDelivery := chatDeliveryHttp.NewHandler(chatUse, zap)

	sessionRepo := sessionRepository.NewRepo(db)
	sessionUse := sessionUsecase.NewUsecase(sessionRepo)

	m := middleware.NewMiddleware(sessionUse, zap)

	r.HandleFunc("/api/chat/users", chatDelivery.GetUsers).Methods("GET")
	r.HandleFunc("/api/chat/messages/{id:[0-9]+}", chatDelivery.GetMessages).Methods("GET")
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		delivery.ServeWs(hub, w, r)
	})

	r.Use(m.AccessLogMiddleware)
	//r.Use(m.CORSMiddleware)
	r.Use(m.AuthMiddleware)

	/*err = http.ListenAndServe(":8002", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
	err = http.ListenAndServe(":8002", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
