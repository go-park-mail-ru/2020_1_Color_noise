package main

import (
	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"
	"2020_1_Color_noise/internal/pkg/middleware"
	"2020_1_Color_noise/internal/pkg/proto/session"
	"google.golang.org/grpc"
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

	db := database.NewPgxDB()
	if err := db.Open(c); err != nil {
		panic(err)
	}

	grcpSessionConn, err := grpc.Dial(
		"auth:8000",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to sessionService")
	}
	defer grcpSessionConn.Close()

	sessManager := session.NewAuthSeviceClient(grcpSessionConn)

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

	m := middleware.NewMiddleware(sessManager, zap)

	r.HandleFunc("/api/chat/users", chatDelivery.GetUsers).Methods("GET")
	r.HandleFunc("/api/chat/messages/{id:[0-9]+}", chatDelivery.GetMessages).Methods("GET")
	r.HandleFunc("/api/chat/stickers", chatDelivery.GetStickers).Methods("GET")
	r.HandleFunc("/api/chat/ws", func(w http.ResponseWriter, r *http.Request) {

		delivery.ServeWs(hub, zap, chatUse, w, r)
	})
	r.Use(m.PanicMiddleware)
	r.Use(m.AccessLogMiddleware)
	//r.Use(m.CORSMiddleware)
	r.Use(m.AuthMiddleware)

	/*err = http.ListenAndServe(":8002", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
