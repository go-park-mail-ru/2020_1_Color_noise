package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/board"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/response"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	boardUsecase board.IUsecase
}

func NewHandler(usecase board.IUsecase) *Handler {
	return &Handler{
		boardUsecase: usecase,
	}
}

func (bh *Handler) Create(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.InputBoard{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		err = error.Wrap(err,"Decoding error during creation board")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Name shouldn't be empty and longer 60 characters. " +
				"Description shouldn't be empty and longer 1000 characters.")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, err := bh.boardUsecase.Create(input, userId)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponseBoard{
		Id: id,
	}

	response.Respond(w, http.StatusCreated, resp)
}

func (bh *Handler) GetBoard(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a board"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	board, err := bh.boardUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponseBoard{
		Id:          board.Id,
		UserId:      board.UserId,
		Name:        board.Name,
		Description: board.Description,
		Pins:        board.Pins,
	}

	response.Respond(w, http.StatusOK, resp)
}

func (bh *Handler) Fetch(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id when in during getting a boards for user"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := bh.boardUsecase.GetByUserId(uint(userId), start, limit)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := make([]models.ResponseBoard, limit)

	for i, board := range pins {
		resp[i] = models.ResponseBoard{
			Id:          board.Id,
			UserId:      board.UserId,
			Name:        board.Name,
			Description: board.Description,
			Pins:        board.Pins,
		}
	}

	response.Respond(w, http.StatusOK, resp)
}

/*

func (ph *PinHandler) UpdatePin(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	oldPin, err := ph.pinUsecase.Get(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		result.Status = "500"
		body["error"] = "Internal error"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	if userId != oldPin.UserId {
		result.Status = "403"
		body["error"] = "No access"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}


	var pin = &models.Pin{
		Name:    r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	err = ph.pinUsecase.Update(uint(id), pin)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	result.Status = "200"
	json.NewEncoder(w).Encode(result)
}

func (ph *PinHandler) DeletePin(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		result.Status = "400"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	pin, err := ph.pinUsecase.Get(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		result.Status = "500"
		body["error"] = "Internal error"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	if userId != pin.UserId {
		result.Status = "403"
		body["error"] = "No access"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	err = ph.pinUsecase.Delete(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}
	result.Status = "200"
	json.NewEncoder(w).Encode(result)
}

*/