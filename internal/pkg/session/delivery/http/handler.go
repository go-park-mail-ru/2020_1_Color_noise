package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pinterest/internal/models"
	"pinterest/internal/pkg/error"
	"pinterest/internal/pkg/response"
	"pinterest/internal/pkg/session"
	"pinterest/internal/pkg/user"
	"time"
)

type Handler struct {
	sessionUsecase  session.IUsecase
	userUsecase  user.IUsecase
}

func NewHandler(sessionUsecase session.IUsecase, userUsecase user.IUsecase) *Handler {
	return &Handler{
		sessionUsecase: sessionUsecase,
		userUsecase: userUsecase,
	}
}

func (sh *Handler) Login(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == true {
		fmt.Println("HHHHHHHHH")
		response.Respond(w, http.StatusOK, map[string]string {
			"message": "Ok",
		})
		return
	}

	input := &models.SignUpInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		err := error.Wrap(err,"Decoding error during login")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	user, err := sh.userUsecase.GetByLogin(input.Login)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	if err = sh.userUsecase.ComparePassword(user, input.Password); err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	session, err := sh.sessionUsecase.CreateSession(user.Id)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   session.Cookie,
		Expires: time.Now().Add(1000 * time.Hour),
		HttpOnly: true,
		Domain: r.Host,
	}

	token := &http.Cookie{
		Name:    "csrf_token",
		Value:   session.Token,
		Expires: time.Now().Add(5 * time.Hour),
		Domain: r.Host,
	}

	http.SetCookie(w, cookie)
	http.SetCookie(w, token)
	response.Respond(w, http.StatusOK, map[string]string {
		"message": "Ok",
	})
}

func (sh *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		err := error.Wrap(err,"Received bad cookie from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}
	/*
	token, err := r.Cookie("token")
	if err != nil {
		err := error.Wrap(err,"Received bad token from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}
	*/
	err = sh.sessionUsecase.Delete(cookie.Value)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	//token.Expires = time.Now().AddDate(0, 0, -1)

	http.SetCookie(w, cookie)
	//http.SetCookie(w, token)

	response.Respond(w, http.StatusOK, map[string]string {
		"message": "Ok",
	})
}