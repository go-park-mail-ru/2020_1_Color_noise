package http

import (
	"bytes"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"io"
	"net/http"
	"pinterest/internal/models"
	"pinterest/internal/pkg/error"
	"pinterest/internal/pkg/response"
	"pinterest/internal/pkg/session"
	"pinterest/internal/pkg/user"
	"time"
)

type Result struct {
	Status string      `json:"status"`
	Body interface{} `json:"body,omitempty"`
}

type Handler struct {
	userUsecase    user.IUsecase
	sessionUsecase session.IUsecase
}

func NewHandler(usecase user.IUsecase, sessionUsecase session.IUsecase) *Handler {
	return &Handler{
		userUsecase: usecase,
		sessionUsecase: sessionUsecase,
	}
}

func (ud *Handler) Create(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == true {
		response.Respond(w, http.StatusOK, map[string]string {
			"message": "Ok",
		})
		return
	}
	input := &models.SignUpInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err = error.Wrap(err,"Decoding error during creation user")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", reqId),
			"Password should be longer than 6 characters and shorter 100. " +
			"Login should be letters and numbers. " +
			"Email should be like hello@example.com")
		error.ErrorHandler(w, err)
		return
	}

	id, err := ud.userUsecase.Create(input)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	session, err := ud.sessionUsecase.CreateSession(id)
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
	response.Respond(w, http.StatusCreated, map[string]string {
		"message": "Ok",
	})
}

func (ud *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, err)
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	user, err := ud.userUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := models.ResponseSettingsUser{
		Email:  user.Email,
		Login:  user.Login,
		About:  user.About,
		Avatar: user.Avatar,
	}

	response.Respond(w, http.StatusOK, resp)
}

func (ud *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.UpdateInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err := error.Wrap(err,"Decoding error during updating user")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ud.userUsecase.Update(id, input)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string {
		"message": "Ok",
	})
}

func (ud *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := map[string]string{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err := error.Wrap(err, "Decoding error during updating password")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	password, ok := input["password"];
	if !ok {
		err := error.BadPassword.New("Enter your password")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ud.userUsecase.UpdatePassword(id, password)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string {
		"message": "Ok",
	})
}

func (ud *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		err := error.Wrap(err, "Decoding error during updating password")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	file, _, err  := r.FormFile("image")
	if err != nil {
		err := error.Wrap(err, "Reading image from form error")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		err := error.Wrap(err, "Coping byte form error")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	address, err := ud.userUsecase.UpdateAvatar(id, buffer)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusCreated, map[string]string{
		"image": address,
	})
}
/*
func (ud *Handler) Delete(w http.ResponseWriter, r *http.Request) {
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

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		result.Status = "500"
		body["error"] = "Invalid id"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	err := ud.userUsecase.Delete(uint(id))
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
*/


