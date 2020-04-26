package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/chat"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
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

func (ph *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch users for chat: user is unauthorized")
		error.ErrorHandler(w, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting users for chat"), "Bad id")
		error.ErrorHandler(w, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	users, err := ph.usecase.GetUsers(uint(userId), start, limit)
	if err != nil {
		error.ErrorHandler(w, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
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

func (ph *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch messages for chat: user is unauthorized")
		error.ErrorHandler(w, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting messages for chat"), "Bad id")
		error.ErrorHandler(w, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	messages, err := ph.usecase.GetMessages(uint(userId), start, limit)
	if err != nil {
		error.ErrorHandler(w, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseMessages, 0)

	for _, message := range messages {
		resp = append(resp, models.ResponseMessages{
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
			Message:   message.Message,
			CreatedAt: message.CreatedAt,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}