package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"2020_1_Color_noise/internal/pkg/session"
	"2020_1_Color_noise/internal/pkg/user"
	"bytes"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
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
			"Login should be letters and numbers, and shorter than 20 characters " +
			"Email should be like hello@example.com and shorter than 50 characters.")
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

	resp := models.ResponseUser{
		Email:  	   user.Email,
		Login:  	   user.Login,
		About:  	   user.About,
		Avatar: 	   user.Avatar,
		Pins:   	   user.Pins,
		Desks:         user.Desks,
		Subscribers:   user.Subscribers,
		Subscriptions: user.Subscriptions,
	}

	response.Respond(w, http.StatusOK, resp)
}

func (ud *Handler) GetOtherUser(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting user"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	user, err := ud.userUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := models.ResponseUser{
		Login:  	   user.Login,
		About:  	   user.About,
		Avatar: 	   user.Avatar,
		Pins:   	   user.Pins,
		Desks:         user.Desks,
		Subscribers:   user.Subscribers,
		Subscriptions: user.Subscriptions,
	}

	response.Respond(w, http.StatusOK, resp)
}

func (ud *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.UpdateProfileInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err := error.Wrap(err,"Decoding error during updating profile user")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", reqId),
				"Login should be letters and numbers, shorter than 20 characters " +
				"Email should be like hello@example.com and shorter than 50 characters")
		error.ErrorHandler(w, err)
		return
	}

	err = ud.userUsecase.UpdateProfile(id, input)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string {
		"message": "Ok",
	})
}

func (ud *Handler) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.UpdateDescriptionInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err := error.Wrap(err,"Decoding error during updating description user")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", reqId),
			"Description should be shorter than 1000 characters.")
		error.ErrorHandler(w, err)
		return
	}

	err = ud.userUsecase.UpdateDescription(id, input)
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

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.UpdatePasswordInput{}

	err := json.NewDecoder(r.Body).Decode(input)
	if err != nil {
		err := error.Wrap(err, "Decoding error during updating password")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", reqId),
			"Password should be longer than 6 characters and shorter 100. ")
		error.ErrorHandler(w, err)
		return
	}

	err = ud.userUsecase.UpdatePassword(id, input)
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

func (ud *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	subId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during following"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ud.userUsecase.Follow(id, uint(subId))
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusCreated, map[string]string {
		"message": "Ok",
	})
}

func (ud *Handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	subId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during unfollowing"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ud.userUsecase.Unfollow(id, uint(subId))
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string {
		"message": "Ok",
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



