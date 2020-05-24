package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/list"
	"2020_1_Color_noise/internal/pkg/response"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase list.IUsecase
	logger  *zap.SugaredLogger
}

func NewHandler(usecase list.IUsecase, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		usecase: usecase,
		logger:  logger,
	}
}

func (lh *Handler) GetMainList(w http.ResponseWriter, r *http.Request) {

	reqId := r.Context().Value("ReqId")

	/*isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch boards: user is unauthorized")
		error.ErrorHandler(w, r, error.Wrapf(err, "request id: %s", reqId))
		return
	}*/

	start, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		start = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := lh.usecase.GetMainList(start, limit)
	if err != nil {
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponsePin, 0)

	for _, pin := range pins {
		resp = append(resp, models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			User:        &models.ResponseUser{
				/*Id: pin.User.Id,
				Login: pin.User.Login,
				About: pin.User.About,
				Avatar: pin.User.Avatar,
				Subscriptions: pin.User.Subscriptions,
				Subscribers: pin.User.Subscribers,*/
			},
			Name:        pin.Name,
			Description: pin.Description,
			Image:       pin.Image,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}

func (lh *Handler) GetSubList(w http.ResponseWriter, r *http.Request) {

	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("GetSubList: user is unauthorized")
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		start = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := lh.usecase.GetSubList(id, start, limit)
	if err != nil {
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponsePin, 0)

	for _, pin := range pins {
		resp = append(resp, models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			User:        &models.ResponseUser{
				Id: pin.User.Id,
				Login: pin.User.Login,
				About: pin.User.About,
				Avatar: pin.User.Avatar,
				Subscriptions: pin.User.Subscriptions,
				Subscribers: pin.User.Subscribers,
			},
			Name:        pin.Name,
			Description: pin.Description,
			Image:       pin.Image,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}

func (lh *Handler) GetRecommendationList(w http.ResponseWriter, r *http.Request) {

	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("GetRecommendationList: user is unauthorized")
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		start = 0
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := lh.usecase.GetRecommendationList(id, start, limit)
	if err != nil {
		error.ErrorHandler(w, r, lh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponsePin, 0)

	for _, pin := range pins {
		resp = append(resp, models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			User:        &models.ResponseUser{
				Id: pin.User.Id,
				Login: pin.User.Login,
				About: pin.User.About,
				Avatar: pin.User.Avatar,
				Subscriptions: pin.User.Subscriptions,
				Subscribers: pin.User.Subscribers,
			},
			Name:        pin.Name,
			Description: pin.Description,
			Image:       pin.Image,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}
