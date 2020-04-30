package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/chat"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/metric"
	"2020_1_Color_noise/internal/pkg/response"
	"fmt"
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

func (ch *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	path := ""
	metric.Increase()
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch users for chat: user is unauthorized")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	users, err := ch.usecase.GetUsers(uint(userId), start, limit)
	if err != nil {
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
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

	response.Respond(w, http.StatusOK, resp, path)
}

func (ch *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	path := ""
	metric.Increase()
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch messages for chat: user is unauthorized")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	vars := mux.Vars(r)
	otherId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting messages for chat"), "Bad id")
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	messages, err := ch.usecase.GetMessages(userId, uint(otherId), start, limit)
	if err != nil {
		error.ErrorHandler(w, ch.logger, reqId, error.Wrapf(err, "request id: %s", reqId), path)
		return
	}

	fmt.Println("messages after usecase: ", messages)

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
			Message:   message.Message,
			CreatedAt: message.CreatedAt,
		})
	}

	fmt.Println("messages in delivery: ", resp)

	response.Respond(w, http.StatusOK, resp, path)
}
