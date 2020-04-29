package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"2020_1_Color_noise/internal/pkg/session"
	"2020_1_Color_noise/internal/pkg/user"
	"fmt"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Handler struct {
	sessionUsecase session.IUsecase
	userUsecase    user.IUsecase
	logger         *zap.SugaredLogger
}

func NewHandler(sessionUsecase session.IUsecase, userUsecase user.IUsecase, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		sessionUsecase: sessionUsecase,
		userUsecase:    userUsecase,
		logger:			logger,
	}
}

func (sh *Handler) Login(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth == true {
		response.Respond(w, http.StatusOK, map[string]string{
			"message": "Ok",
		})
		return
	}

	input := &models.SignUpInput{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during login"), "Wrong body of request")
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	user, err := sh.userUsecase.GetByLogin(input.Login)
	if err != nil {
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	if err = sh.userUsecase.ComparePassword(user, input.Password); err != nil {
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	session, err := sh.sessionUsecase.CreateSession(user.Id)
	if err != nil {
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    session.Cookie,
		Expires:  time.Now().Add(1000 * time.Hour),
		HttpOnly: true,
		//Domain:   r.Host,
	}

	resp := models.ResponseUser{
		Id:            user.Id,
		Email:         user.Email,
		Login:         user.Login,
		About:         user.About,
		Avatar:        user.Avatar,
		Subscribers:   user.Subscribers,
		Subscriptions: user.Subscriptions,
	}

	token := &http.Cookie{
		Name:    "csrf_token",
		Value:   session.Token,
		Expires: time.Now().Add(5 * time.Hour),
		//Domain:  r.Host,
	}

	http.SetCookie(w, cookie)
	http.SetCookie(w, token)
	fmt.Println(w)

	response.Respond(w, http.StatusOK, resp)
}

func (sh *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Logout session: user is unauthorized")
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		err := error.Wrap(err, "Received bad cookie from context")
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
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
		error.ErrorHandler(w, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	//token.Expires = time.Now().AddDate(0, 0, -1)

	http.SetCookie(w, cookie)
	fmt.Println(w)
	//http.SetCookie(w, token)

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	})
}
