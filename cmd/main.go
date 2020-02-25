package main

import (
	httpDeliverPin "pinterest/pkg/pin/delivery/http"
	repoPin "pinterest/pkg/pin/repository"
	usecasePin "pinterest/pkg/pin/usecase"
	//"fmt"
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

/*
var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type MyHandler struct {
	sessions map[string]uint
	users    map[string]*User
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		sessions: make(map[string]uint, 10),
		users: map[string]*User{
			"rvasily": {1, "rvasily", "love"},
		},
	}
}

// http://127.0.0.1:8080/login?login=rvsily&password=love

func (api *MyHandler) Login(w http.ResponseWriter, r *http.Request) {

	user, ok := api.users[r.FormValue("login")]
	//fmt.Println(user, api.users[r.FormValue("login")], api.users["rvasily"],r.FormValue("login"))
	if !ok {
		http.Error(w, `no user`, 404)
		return
	}

	if user.Password != r.FormValue("password") {
		http.Error(w, `bad pass`, 400)
		return
	}

	SID := RandStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
	//w.Write([]byte(SID))

}

func (api *MyHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	user, ok := api.users[r.FormValue("login")]
	//fmt.Println(user, api.users[r.FormValue("login")], api.users["rvasily"],r.FormValue("login"))
	if !ok {
		http.Error(w, `no user`, 404)
		return
	}

	if user.Password != r.FormValue("password") {
		http.Error(w, `bad pass`, 400)
		return
	}

	SID := RandStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
	//w.Write([]byte(SID))

}

func (api *MyHandler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no sess`, 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, `no sess`, 401)
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (api *MyHandler) Root(w http.ResponseWriter, r *http.Request) {
	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
 */

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
	//api := NewMyHandler()
	/*sessionService := delivery.NewSessionHTTP()
	sessionUseCase := usecase.NewSessionUseCase()
	sessionService.SessionUseCase = sessionUseCase

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/session/{id}", sessionService.GetSession).Methods("GET")*/
	//r.PathPrefix("/sample-2020-lesson-2/public/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	//r.HandleFunc("/", api.Root)
	r.HandleFunc("/login", sessionDelivery.Login).Methods("POST")
	r.HandleFunc("/logout", sessionDelivery.Logout)
	r.HandleFunc("/signup", userDelivery.AddUser).Methods("POST")
	r.HandleFunc("/profile/{id}", userDelivery.GetUser).Methods("GET")
	r.HandleFunc("/profile/{id}", userDelivery.UpdateUser).Methods("POST")
	r.HandleFunc("/profile/{id}", userDelivery.DeleteUser).Methods("DELETE")
	r.HandleFunc("/pin", pinDelivery.Add).Methods("POST")
	r.HandleFunc("/pin/{id}", pinDelivery.GetPin).Methods("GET")
	r.HandleFunc("/pin/{id}", pinDelivery.UpdatePin).Methods("PUT")
	r.HandleFunc("/pin/{id}", pinDelivery.DeletePin).Methods("DELETE")
	r.Use(m.AuthMiddleware)
	r.Use(m.CORSMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
