package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/notifications"
	"2020_1_Color_noise/internal/pkg/response"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase notifications.IUsecase
}

func NewHandler(usecase notifications.IUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (nh *Handler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("GetNotifacations: user is unauthorized")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	notifications, err := nh.usecase.GetNotifications(id, start, limit)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseNotification, 0)

	for _, notification := range notifications {
		resp = append(resp, models.ResponseNotification{
			User:	 notification.User,
			Message: notification.Message,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}
