package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	authService "2020_1_Color_noise/internal/pkg/proto/session"
	userService "2020_1_Color_noise/internal/pkg/proto/user"
	"2020_1_Color_noise/internal/pkg/response"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type Handler struct {
	as     authService.AuthSeviceClient
	us     userService.UserServiceClient
	logger *zap.SugaredLogger
}

func NewHandler(as authService.AuthSeviceClient, us userService.UserServiceClient, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		as,
		us,
		logger,
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
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}
	/*
		user, err := sh.as.GetByLogin(input.Login)
		if err != nil {
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}

		if err = sh.userUsecase.ComparePassword(user, input.Password); err != nil {
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}

		session, err := sh.sessionUsecase.Create(user.Id)
		if err != nil {
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}
	*/
	u, err := sh.us.GetByLogin(r.Context(), &userService.Login{Login: input.Login})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	sess, err := sh.as.Login(r.Context(), &authService.SignIn{
		User: &authService.User{
			Id:                u.Id,
			Email:             u.Email,
			Login:             u.Login,
			EncryptedPassword: u.EncryptedPassword,
			About:             u.About,
			Avatar:            u.Avatar,
			Subscribers:       u.Subscribers,
			Subscriptions:     u.Subscriptions,
		},
		Password: input.Password,
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.Cookie,
		Expires:  time.Now().Add(1000 * time.Hour),
		HttpOnly: true,
		Domain:   r.Host,
	}

	token := &http.Cookie{
		Name:    "csrf_token",
		Value:   sess.Token,
		Expires: time.Now().Add(1000 * time.Hour),
		Domain:  r.Host,
	}

	http.SetCookie(w, cookie)
	http.SetCookie(w, token)

	resp := models.ResponseUser{
		Id:            uint(u.Id),
		Email:         u.Email,
		Login:         u.Login,
		About:         u.About,
		Avatar:        u.Avatar,
		Subscribers:   int(u.Subscribers),
		Subscriptions: int(u.Subscriptions),
	}

	response.Respond(w, http.StatusOK, resp)
}

func (sh *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Logout session: user is unauthorized")
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		err := error.Wrap(err, "Received bad cookie from context")
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}
	/*
		token, err := r.Cookie("token")
		if err != nil {
			err := error.Wrap(err,"Received bad token from context")
			error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}
	*/

	_, err = sh.as.Delete(r.Context(), &authService.Cookie{
		Cookie: cookie.Value,
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, sh.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	//token.Expires = time.Now().AddDate(0, 0, -1)

	http.SetCookie(w, cookie)
	//http.SetCookie(w, token)

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	})
}
