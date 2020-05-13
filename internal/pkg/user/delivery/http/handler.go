package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	authService "2020_1_Color_noise/internal/pkg/proto/session"
	userService "2020_1_Color_noise/internal/pkg/proto/user"
	"2020_1_Color_noise/internal/pkg/response"
	"bytes"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	us     userService.UserServiceClient
	as     authService.AuthSeviceClient
	logger *zap.SugaredLogger
}

func NewHandler(us userService.UserServiceClient, as authService.AuthSeviceClient, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		us,
		as,
		logger,
	}
}

func (ud *Handler) Create(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth == true {
		response.Respond(w, http.StatusOK, map[string]string{
			"message": "Ok",
		} )
		return
	}

	input := &models.SignUpInput{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during creation user"), "Wrong body of request")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	u, err := ud.us.Create(r.Context(), &userService.SignUp{
		Email: input.Email,
		Login: input.Login,
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
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	sess, err := ud.as.Create(r.Context(), &authService.UserID{Id: int64(u.Id)})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
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
			Expires: time.Now().Add(5 * time.Hour),
			Domain:  r.Host,
	}


	resp := models.ResponseUser{
		Id:            uint(u.Id),
		Login:         u.Login,
		About:         u.About,
		Avatar:        u.Avatar,
	}

	http.SetCookie(w, cookie)
	http.SetCookie(w, token)

	response.Respond(w, http.StatusCreated, resp )
}

func (ud *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf( err,"request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	us, err := ud.us.GetById(r.Context(), &userService.UserID{Id: int64(id)})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	resp := models.ResponseUser{
		Id:            uint(us.Id),
		Email:         us.Email,
		Login:         us.Login,
		About:         us.About,
		Avatar:        us.Avatar,
		Subscribers:   int(us.Subscribers),
		Subscriptions: int(us.Subscriptions),
	}

	response.Respond(w, http.StatusOK, resp )
}

func (ud *Handler) GetOtherUser(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get other user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting user"), "Bad id")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	us, err := ud.us.GetById(r.Context(), &userService.UserID{Id: int64(id)})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	resp := models.ResponseUser{
		Id:            uint(us.Id),
		Login:         us.Login,
		About:         us.About,
		Avatar:        us.Avatar,
		Subscribers:   int(us.Subscribers),
		Subscriptions: int(us.Subscriptions),
	}

	response.Respond(w, http.StatusOK, resp )
}

func (ud *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Update user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	input := &models.UpdateProfileInput{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during updating profile user"), "Wrong body of request")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	_, err = ud.us.UpdateProfile(r.Context(), &userService.Profile{
		Id: &userService.UserID{Id: int64(id)},
		Input: &userService.UpdateProfileInput{Login: input.Login,
			Email: input.Email,
			},
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	} )
}

func (ud *Handler) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Update description of user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	input := &models.UpdateDescriptionInput{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during updating description user"), "Wrong body of request")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	_, err = ud.us.UpdateDescription(r.Context(), &userService.Description{
		Id: &userService.UserID{Id: int64(id)},
		Description: input.Description,
			})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	} )
}

func (ud *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Update password of user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	input := &models.UpdatePasswordInput{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during updating password user"), "Wrong body of request")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	_, err = ud.us.UpdatePassword(r.Context(), &userService.Password{
		Id: &userService.UserID{Id: int64(id)},
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
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	} )
}

func (ud *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Upload avatar of user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		err := error.Wrap(err, "Decoding error during updating password")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		err := error.Wrap(err, "Reading image from form error")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		err := error.Wrap(err, "Coping byte form error")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	address, err := ud.us.UpdateAvatar(r.Context(), &userService.Avatar{
		Id: &userService.UserID{Id: int64(id)},
		Avatar: buffer.Bytes(),
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusCreated, map[string]string{
		"image": address.Avatar,
	} )
}

func (ud *Handler) Follow(w http.ResponseWriter, r *http.Request) {


	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Following user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	vars := mux.Vars(r)
	subId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during following"), "Bad id")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	if id == uint(subId) {
		err = error.WithMessage(error.BadRequest.New("Bad id in during following user"),
			"Your id and following id shoudn't match")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	_, err = ud.us.Follow(r.Context(), &userService.Following{
		Id: &userService.UserID{Id: int64(id)},
		SubId: &userService.UserID{Id: int64(subId)},
			})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusCreated, map[string]string{
		"message": "Ok",
	} )
}

func (ud *Handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Unfollowing  user: user is unauthorized")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	vars := mux.Vars(r)
	subId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during unfollowing"), "Bad id")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	if id == uint(subId) {
		err = error.WithMessage(error.BadRequest.New("Bad id in during unfollowing user"),
			"Your id and unfollowing id shoudn't match")
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	_, err = ud.us.Unfollow(r.Context(), &userService.Following{
		Id: &userService.UserID{Id: int64(id)},
		SubId: &userService.UserID{Id: int64(subId)},
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, ud.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	} )
}

func (uh *Handler) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get subscribers of user: user is unauthorized")
		error.ErrorHandler(w, r, uh.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting subscribers for user"), "Bad id")
		error.ErrorHandler(w, r, uh.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	users, err := uh.us.GetSubscribers(r.Context(), &userService.Sub{
		Id: &userService.UserID{Id: int64(id)},
		Start: int64(start),
		Limit: int64(limit),
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, uh.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseUser, 0)

	for _, user := range users.Users {
		resp = append(resp, models.ResponseUser{
			Id:            uint(user.Id),
			Login:         user.Login,
			About:         user.About,
			Avatar:        user.Avatar,
			Subscribers:   int(user.Subscribers),
			Subscriptions: int(user.Subscriptions),
		})
	}

	response.Respond(w, http.StatusOK, resp )
}

func (uh *Handler) GetSubscribtions(w http.ResponseWriter, r *http.Request) {
	  
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get subscribtions of user: user is unauthorized")
		error.ErrorHandler(w, r, uh.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting subscribtions for user"), "Bad id")
		error.ErrorHandler(w, r, uh.logger, reqId, error.Wrapf(err, "request id: %s", reqId) )
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	users, err := uh.us.GetSubscriptions(r.Context(), &userService.Sub{
		Id: &userService.UserID{Id: int64(id)},
		Start: int64(start),
		Limit: int64(limit),
	})
	if err != nil {
		e := error.NoType
		errStatus, ok := status.FromError(err)
		msg := "Unknown GRPC error"
		if ok == true {
			e = error.Cast(int(errStatus.Code()))
			msg = errStatus.Message()
		}
		error.ErrorHandler(w, r, uh.logger, reqId, error.Wrapf(e.New(msg), "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseUser, 0)

	for _, user := range users.Users {
		resp = append(resp, models.ResponseUser{
			Id:            uint(user.Id),
			Login:         user.Login,
			About:         user.About,
			Avatar:        user.Avatar,
			Subscribers:   int(user.Subscribers),
			Subscriptions: int(user.Subscriptions),
		})
	}

	response.Respond(w, http.StatusOK, resp )
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
