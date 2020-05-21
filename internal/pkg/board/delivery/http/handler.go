package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	boardUsecase board.IUsecase
	logger       *zap.SugaredLogger
}

func NewHandler(usecase board.IUsecase, logger *zap.SugaredLogger) *Handler {
	return &Handler{
		boardUsecase: usecase,
		logger:       logger,
	}
}

func (bh *Handler) Create(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Create board: user is unauthorized")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.InputBoard{}

	err := easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during creation board"), "Wrong body of request")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Name shouldn't be empty and longer 60 characters. "+
				"Description shouldn't be empty and longer 1000 characters.")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, err := bh.boardUsecase.Create(input, userId)
	if err != nil {
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponseBoard{
		Id: id,
	}

	response.Respond(w, http.StatusCreated, resp)
}

func (bh *Handler) GetBoard(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	/*
		isAuth := r.Context().Value("IsAuth")
		if isAuth != true {
			err := error.Unauthorized.New("Get board: user is unauthorized")
			error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}

	*/

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a board"), "Bad id")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	board, err := bh.boardUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	respPin := make([]*models.ResponsePin, 0)

	for _, pin := range board.Pins {
		respPin = append(respPin, &models.ResponsePin{
			Id:          pin.Id,
			BoardId:     pin.BoardId,
			UserId:      pin.UserId,
			Name:        pin.Name,
			Description: pin.Description,
			Image:       pin.Image,
		})
	}

	resp := &models.ResponseBoard{
		Id:          board.Id,
		UserId:      board.UserId,
		Name:        board.Name,
		Description: board.Description,
		Pins:        respPin,
		/*LastPin:     &models.ResponsePin{
			Id:          board.LastPin.Id,
			BoardId:     board.LastPin.BoardId,
			UserId:      board.LastPin.UserId,
			Name:        board.LastPin.Name,
			Description: board.LastPin.Description,
			Image:       board.LastPin.Image,
		},*/
	}

	response.Respond(w, http.StatusOK, resp)
}

func (bh *Handler) GetNameBoard(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	/*
		isAuth := r.Context().Value("IsAuth")
		if isAuth != true {
			err := error.Unauthorized.New("Get name of board: user is unauthorized")
			error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}

	*/

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a board"), "Bad id")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	board, err := bh.boardUsecase.GetByNameId(uint(id))
	if err != nil {
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponseBoard{
		Id:          board.Id,
		UserId:      board.UserId,
		Name:        board.Name,
		Description: board.Description,
		LastPin: &models.ResponsePin{
			Id:          board.LastPin.Id,
			BoardId:     board.LastPin.BoardId,
			UserId:      board.LastPin.UserId,
			Name:        board.LastPin.Name,
			Description: board.LastPin.Description,
			Image:       board.LastPin.Image,
		},
	}

	response.Respond(w, http.StatusOK, resp)
}

func (bh *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	/*
		isAuth := r.Context().Value("IsAuth")
		if isAuth != true {
			err := error.Unauthorized.New("Fetch boards: user is unauthorized")
			error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
			return
		}
	*/

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id when in during getting boards for user"), "Bad id")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	boards, err := bh.boardUsecase.GetByUserId(uint(id), start, limit)
	if err != nil {
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseBoard, 0)

	for _, board := range boards {
		resp = append(resp, models.ResponseBoard{
			Id:          board.Id,
			UserId:      board.UserId,
			Name:        board.Name,
			Description: board.Description,
			LastPin: &models.ResponsePin{
				Id:          board.LastPin.Id,
				BoardId:     board.LastPin.BoardId,
				UserId:      board.LastPin.UserId,
				Name:        board.LastPin.Name,
				Description: board.LastPin.Description,
				Image:       board.LastPin.Image,
			},
		})
	}

	response.Respond(w, http.StatusOK, resp)
}

func (bh *Handler) Update(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Update board: user is unauthorized")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id when in during getting boards for user"), "Bad id")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.InputBoard{}

	err = easyjson.UnmarshalFromReader(r.Body, input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Decoding error during updating board"), "Wrong body of request")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Name shouldn't be empty and longer 60 characters. "+
				"Description shouldn't be empty and longer 1000 characters.")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = bh.boardUsecase.Update(input, uint(id), userId)
	if err != nil {
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	})
}

func (bh *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	reqId := r.Context().Value("ReqId")

	isAuth := r.Context().Value("IsAuth")
	if isAuth != true {
		err := error.Unauthorized.New("Get board: user is unauthorized")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a board"), "Bad id")
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	err = bh.boardUsecase.Delete(uint(id), userId)
	if err != nil {
		error.ErrorHandler(w, r, bh.logger, reqId, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	response.Respond(w, http.StatusOK, map[string]string{
		"message": "Ok",
	})
}
