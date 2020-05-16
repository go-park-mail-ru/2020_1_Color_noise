package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/pin"
	"2020_1_Color_noise/internal/pkg/response"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	pinUsecase pin.IUsecase
	logger  *zap.SugaredLogger
}

func NewHandler(usecase pin.IUsecase, logger  *zap.SugaredLogger) *Handler {
	return &Handler{
		pinUsecase: usecase,
		logger:     logger,
	}
}

func (ph *Handler) Create(w http.ResponseWriter, r *http.Request) {
	
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Create pin: user is unauthorized")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.InputPin{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during creation pin"), "Wrong body of request")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Name shouldn't be empty and longer 60 characters. "+
				"Description shouldn't be empty and longer 1000 characters. "+
				"Image should be base64")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, err := ph.pinUsecase.Create(input, userId)
	if err != nil {
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponsePin{
		Id: id,
	}

	response.Respond(w, http.StatusCreated, resp)
}

func (ph *Handler) GetPin(w http.ResponseWriter, r *http.Request) {
	
	reqId := r.Context().Value("ReqId")

	/*
	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get pin: user is unauthorized")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}
	*/

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a pin"), "Bad id")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	pin, err := ph.pinUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponsePin{
		Id:          pin.Id,
		BoardId:     pin.BoardId,
		UserId:      pin.UserId,
		Name:        pin.Name,
		Description: pin.Description,
		Image:       pin.Image,
	}

	response.Respond(w, http.StatusOK, resp)
}

func (ph *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Fetch pin: user is unauthorized")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting pins for user"), "Bad id")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := ph.pinUsecase.GetByUserId(uint(userId), start, limit)
	if err != nil {
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
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

func (ph *Handler) Update(w http.ResponseWriter, r *http.Request) {
	
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Update pin: user is unauthorized")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	pinId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during updating pin"), "Bad id")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.UpdatePin{}

	err = easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during creation pin"), "Wrong body of request")
		err = error.Wrap(err, "Decoding error during updating pin")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Name shouldn't be empty and longer 60 characters. "+
				"Description shouldn't be empty and longer 1000 characters.")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ph.pinUsecase.Update(input, uint(pinId), userId)
	if err != nil {
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	})
}

func (ph *Handler) DeletePin(w http.ResponseWriter, r *http.Request) {
	
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Delete pin: user is unauthorized")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	pinId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during deleting pin"), "Bad id")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ph.pinUsecase.Delete(uint(pinId), userId)
	if err != nil {
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	})
}

func (ph *Handler) Save(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Save pin: user is unauthorized")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	pinId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during saving a pin"), "Bad id")
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = ph.pinUsecase.Save(uint(pinId), userId)
	if err != nil {
		error.ErrorHandler(w, r, ph.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusCreated, map[string]string{
		"message": "Ok",
	})
}
