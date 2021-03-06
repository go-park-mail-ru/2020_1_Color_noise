package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/chat"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase chat.IUsecase
	logger  *zap.SugaredLogger
}

func NewHandler(usecase chat.IUsecase, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		usecase: usecase,
		logger:  logger,
	}
}

func (ch *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch users for chat: user is unauthorized")
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, t := strconv.Atoi(r.URL.Query().Get("start"))
	if t != nil {
		t = nil
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	users, err := ch.usecase.GetUsers(uint(userId), start, limit)
	if err != nil {
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseUser, 0)

	for _, user := range users {
		resp = append(resp, models.ResponseUser{
			Id:            user.Id,
			Login:         user.Login,
			About:         user.About,
			Avatar:        user.Avatar,
			Subscribers:   user.Subscribers,
			Subscriptions: user.Subscriptions,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}

func (ch *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch messages for chat: user is unauthorized")
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	otherId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting messages for chat"), "Bad id")
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	messages, err := ch.usecase.GetMessages(userId, uint(otherId), start, limit)
	if err != nil {
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseMessage, 0)

	for _, message := range messages {
		resp = append(resp, models.ResponseMessage{
			SendUser: &models.ResponseUser{
				Id:            message.SendUser.Id,
				Login:         message.SendUser.Login,
				About:         message.SendUser.About,
				Avatar:        message.SendUser.Avatar,
				Subscribers:   message.SendUser.Subscribers,
				Subscriptions: message.SendUser.Subscriptions,
			},
			RecUser: &models.ResponseUser{
				Id:            message.RecUser.Id,
				Login:         message.RecUser.Login,
				About:         message.RecUser.About,
				Avatar:        message.RecUser.Avatar,
				Subscribers:   message.RecUser.Subscribers,
				Subscriptions: message.RecUser.Subscriptions,
			},
			Stickers:  message.Stickers,
			Message:   message.Message,
			CreatedAt: message.CreatedAt,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}

func (ch *Handler) GetStickers(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch messages for chat: user is unauthorized")
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	files, err := ioutil.ReadDir("../static/stickers")
	if err != nil {
		error.ErrorHandler(w, r, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]string, 0)

	for _, file := range files {
		resp = append(resp, "stickers/"+file.Name())
	}

	response.Respond(w, http.StatusOK, resp)
}
