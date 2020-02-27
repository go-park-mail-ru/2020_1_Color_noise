package http

import (
	"encoding/json"
	"fmt"

	//"github.com/gorilla/mux"
	"net/http"
	"pinterest/pkg/models"
	sessionUsecase "pinterest/pkg/session/usecase"
	//"strconv"
	userUsecase "pinterest/pkg/user/usecase"
	"time"
)

type Result struct {
	Status string      `json:"status"`
	Body interface{} `json:"body,omitempty"`
}

type SessionHandler struct {
	sessionUsecase  *sessionUsecase.SessionUsecase
	userUsecase  *userUsecase.UserUsecase
}

func NewSessionHandler(sessionUsecase *sessionUsecase.SessionUsecase, userUsecase *userUsecase.UserUsecase) *SessionHandler {
	return &SessionHandler{
		sessionUsecase: sessionUsecase,
		userUsecase: userUsecase,
	}
}

func (sh *SessionHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := Result{}
	body := map[string]interface{} {}

	if r.Context().Value("isAuth") == true {
		result.Status = "200"
		body["id"] = r.Context().Value("Id")
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	in := &models.User{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	user, err := sh.userUsecase.GetByLogin(in.Login)
	fmt.Println(user, in)
	if err != nil {
		result.Status = "403"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	if !sh.userUsecase.ComparePassword(user, in.Password) {
		result.Status = "403"
		body["error"] = "User not found"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	session, err := sh.sessionUsecase.CreateSession(user.Id)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   session.Cookie,
		Expires: time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Domain: r.Host,
	}
	token := &http.Cookie{
		Name:    "csrf_token",
		Value:   session.Token,
		Expires: time.Now().Add(10 * time.Hour),
		Domain: r.Host,
	}
	result.Status = "200"
	body["id"] = user.Id
	result.Body = body
	http.SetCookie(w, cookie)
	http.SetCookie(w, token)
	json.NewEncoder(w).Encode(result)
}

func (sh *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := Result{}
	body := map[string]interface{} {}

	if r.Context().Value("isAuth") == false {
		result.Status = "403"
		body["error"] = "User not found"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err := sh.sessionUsecase.Delete(cookie.Value)
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	token, _ := r.Cookie("token")
	token.Expires = time.Now().AddDate(0, 0, -1)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	result.Status = "200"
	http.SetCookie(w, cookie)
	http.SetCookie(w, token)
	json.NewEncoder(w).Encode(result)
}
