package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/list"
	"2020_1_Color_noise/internal/pkg/response"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase list.IUsecase
}

func NewHandler(usecase list.IUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (lh *Handler) GetMainList(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	/*isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch boards: user is unauthorized")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}*/

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := lh.usecase.GetMainList(start, limit)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponsePin, 0)

	for _, pin := range pins {
		resp = append(resp, models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			UserId:      pin.UserId,
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
		err := error.Unauthorized.New("Fetch boards: user is unauthorized")
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

	pins, err := lh.usecase.GetSubList(id, start, limit)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponsePin, 0)

	for _, pin := range pins {
		resp = append(resp, models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			UserId:      pin.UserId,
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
		err := error.Unauthorized.New("Fetch boards: user is unauthorized")
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

	pins, err := lh.usecase.GetRecommendationList(id, start, limit)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponsePin, 0)

	for _, pin := range pins {
		resp = append(resp, models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			UserId:      pin.UserId,
			Name:        pin.Name,
			Description: pin.Description,
			Image:       pin.Image,
		})
	}

	response.Respond(w, http.StatusOK, resp)
}
