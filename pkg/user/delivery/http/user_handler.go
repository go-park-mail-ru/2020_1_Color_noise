package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"pinterest/pkg/models"
	sessionUsecase "pinterest/pkg/session/usecase"
	userUsecase "pinterest/pkg/user/usecase"
	"strconv"
	"time"
)

type Result struct {
	Status string      `json:"status"`
	Body interface{} `json:"body,omitempty"`
}

type UserDelivery struct {
	userUsecase  *userUsecase.UserUsecase
	sessionUsecase *sessionUsecase.SessionUsecase
}

func NewUserDelivery(usecase *userUsecase.UserUsecase, sessionUsecase *sessionUsecase.SessionUsecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: usecase,
		sessionUsecase: sessionUsecase,
	}
}

func (ud *UserDelivery) AddUser(w http.ResponseWriter, r *http.Request) {
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

	user := &models.User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}

	id, err := ud.userUsecase.Add(user)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	session, err := ud.sessionUsecase.CreateSession(id)
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
	body["id"] = id
	result.Body = body
	http.SetCookie(w, cookie)
	http.SetCookie(w, token)
	json.NewEncoder(w).Encode(result)
}

func (ud *UserDelivery) GetUser(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		result.Status = "404"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	user, err := ud.userUsecase.GetById(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	result.Status = "200"
	body["user"] = user
	result.Body = body
	json.NewEncoder(w).Encode(result)
}

func (ud *UserDelivery) UpdateUser(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("1111111")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		result.Status = "400"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	fmt.Println("22222222")
	oldUser, err := ud.userUsecase.GetById(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = "Internal"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		result.Status = "500"
		body["error"] = "Internal error"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	if userId != oldUser.Id {
		result.Status = "500"
		body["error"] = "Internal error"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	image := bytes.NewBuffer(nil)
	file, _, err  := r.FormFile("avatar")
	if err == nil {
		io.Copy(image, file)
	}
	fmt.Println(oldUser.Id, "here")
	var user = &models.User{
		Id:         oldUser.Id,
		Email:      r.FormValue("email"),
		Login:      r.FormValue("login"),
		Password:   r.FormValue("password"),
		About:      r.FormValue("about"),
		DataAvatar: image.Bytes(),
	}

	err = ud.userUsecase.Update(user)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	result.Status = "200"
	json.NewEncoder(w).Encode(result)
}

func (ud *UserDelivery) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		result.Status = "400"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	oldUser, err := ud.userUsecase.GetById(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = "Internal"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		result.Status = "500"
		body["error"] = "Internal error"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	if userId != oldUser.Id {
		result.Status = "500"
		body["error"] = "Internal error"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	err = ud.userUsecase.Delete(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	cookie, _ := r.Cookie("session_id")
	ud.sessionUsecase.Delete(cookie.Value)
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	token, _ := r.Cookie("token")
	token.Expires = time.Now().AddDate(0, 0, -1)
	result.Status = "200"
	http.SetCookie(w, cookie)
	http.SetCookie(w, token)
	json.NewEncoder(w).Encode(result)
}

