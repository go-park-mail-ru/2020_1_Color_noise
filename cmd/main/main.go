package main

import (
	boardDeliveryHttp "2020_1_Color_noise/internal/pkg/board/delivery/http"
	boardRepository "2020_1_Color_noise/internal/pkg/board/repository"
	boardUsecase "2020_1_Color_noise/internal/pkg/board/usecase"
	"2020_1_Color_noise/internal/pkg/metric"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"2020_1_Color_noise/internal/pkg/proto/session"
	"2020_1_Color_noise/internal/pkg/proto/user"

	commentDeliveryHttp "2020_1_Color_noise/internal/pkg/comment/delivery/http"
	commentRepository "2020_1_Color_noise/internal/pkg/comment/repository"
	commentUsecase "2020_1_Color_noise/internal/pkg/comment/usecase"

	"2020_1_Color_noise/internal/pkg/config"
	"2020_1_Color_noise/internal/pkg/database"

	listDeliveryHttp "2020_1_Color_noise/internal/pkg/list/delivery/http"
	listRepository "2020_1_Color_noise/internal/pkg/list/repository"
	listUsecase "2020_1_Color_noise/internal/pkg/list/usecase"

	"2020_1_Color_noise/internal/pkg/middleware"

	notificationsDeliveryHttp "2020_1_Color_noise/internal/pkg/notifications/delivery/http"
	notificationsRepository "2020_1_Color_noise/internal/pkg/notifications/repository"
	notificationsUsecase "2020_1_Color_noise/internal/pkg/notifications/usecase"

	pinDeliveryHttp "2020_1_Color_noise/internal/pkg/pin/delivery/http"
	pinRepository "2020_1_Color_noise/internal/pkg/pin/repository"
	pinUsecase "2020_1_Color_noise/internal/pkg/pin/usecase"

	imageUsecase "2020_1_Color_noise/internal/pkg/image/usecase"

	searchHandler "2020_1_Color_noise/internal/pkg/search"

	sessionDeliveryHttp "2020_1_Color_noise/internal/pkg/session/delivery/http"
	userDeliveryHttp "2020_1_Color_noise/internal/pkg/user/delivery/http"
	"go.uber.org/zap"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	metric.Register()
	r := mux.NewRouter()

	c, err := config.GetDBConfing()
	if err != nil {
		panic(err)
	}

	db := database.NewPgxDB()
	if err = db.Open(c); err != nil {
		panic(err)
	}

	grcpSessConn, err := grpc.Dial(
		"auth:8000",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to session grpc")
	}
	defer grcpSessConn.Close()

	sessManager := session.NewAuthSeviceClient(grcpSessConn)

	grcpUserConn, err := grpc.Dial(
		"user:8000",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to user grpc")
	}
	defer grcpSessConn.Close()

	userService := user.NewUserServiceClient(grcpUserConn)

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap := logger.Sugar().With(
		zap.String("mode", "[access_log]"),
		zap.String("logger", "ZAP"),
	)

	sessionDelivery := sessionDeliveryHttp.NewHandler(sessManager, userService, zap)

	boardRepo := boardRepository.NewRepo(db)
	boardUse := boardUsecase.NewUsecase(boardRepo)
	boardDelivery := boardDeliveryHttp.NewHandler(boardUse, zap)

	pinRepo := pinRepository.NewRepo(db)

	imageUse := imageUsecase.NewImageUsecase(pinRepo, boardRepo, userService)
	pinUse := pinUsecase.NewUsecase(pinRepo, boardRepo, imageUse)
	pinDelivery := pinDeliveryHttp.NewHandler(pinUse, zap)

	commentRepo := commentRepository.NewRepo(db)
	commentUse := commentUsecase.NewUsecase(commentRepo)
	commentDelivery := commentDeliveryHttp.NewHandler(commentUse, zap)

	listRepo := listRepository.NewRepo(db)
	listUse := listUsecase.NewUsecase(listRepo)
	listDelivery := listDeliveryHttp.NewHandler(listUse, zap)

	notificationsRepo := notificationsRepository.NewRepo(db)
	notificationsUse := notificationsUsecase.NewUsecase(notificationsRepo)
	notificationsDelivery := notificationsDeliveryHttp.NewHandler(notificationsUse, zap)

	search := searchHandler.NewHandler(commentUse, pinUse, userService, zap)

	userDelivery := userDeliveryHttp.NewHandler(userService, sessManager, boardUse, zap)

	m := middleware.NewMiddleware(sessManager, zap)

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
	r.HandleFunc("/api/support", userDelivery.GetSupport).Methods("GET")
	r.HandleFunc("/api/user/following/{id:[0-9]+}/status", userDelivery.IsFollowed).Methods("GET")
	//r.HandleFunc("/api/user", userDelivery.Delete).Methods("DELETE")

	r.HandleFunc("/api/pin", pinDelivery.Create).Methods("POST")
	r.HandleFunc("/api/pin/saving/{id:[0-9]+}", pinDelivery.Save).Methods("POST")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.GetPin).Methods("GET")
	r.HandleFunc("/api/pin/user/{id:[0-9]+}", pinDelivery.Fetch).Methods("GET")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.Update).Methods("PUT")
	r.HandleFunc("/api/pin/{id:[0-9]+}", pinDelivery.DeletePin).Methods("DELETE")

	r.HandleFunc("/api/board", boardDelivery.Create).Methods("POST")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.Update).Methods("PUT")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.Delete).Methods("DELETE")
	r.HandleFunc("/api/board/{id:[0-9]+}", boardDelivery.GetBoard).Methods("GET")
	r.HandleFunc("/api/board/name/user/{id:[0-9]+}", boardDelivery.GetNameBoard).Methods("GET")
	r.HandleFunc("/api/board/user/{id:[0-9]+}", boardDelivery.Fetch).Methods("GET")

	r.HandleFunc("/api/comment", commentDelivery.Create).Methods("POST")
	r.HandleFunc("/api/comment/{id:[0-9]+}", commentDelivery.GetComment).Methods("GET")
	r.HandleFunc("/api/comment/pin/{id:[0-9]+}", commentDelivery.Fetch).Methods("GET")

	r.HandleFunc("/api/search", search.Search).Methods("GET")

	r.HandleFunc("/api/list", listDelivery.GetMainList).Methods("GET")
	r.HandleFunc("/api/list/sub", listDelivery.GetSubList).Methods("GET")
	r.HandleFunc("/api/list/recommendation", listDelivery.GetRecommendationList).Methods("GET")

	r.HandleFunc("/api/notifications", notificationsDelivery.GetNotifications).Methods("GET")

	r.Handle("/api/metric", promhttp.Handler())
	r.Use(m.PanicMiddleware)
	r.Use(m.AccessLogMiddleware)
	//r.Use(m.CORSMiddleware)
	r.Use(m.AuthMiddleware)

	/*r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))*/

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
