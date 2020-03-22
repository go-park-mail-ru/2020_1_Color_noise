package http

import (
	"2020_1_Color_noise/internal/models"
	"2020_1_Color_noise/internal/pkg/error"
	"2020_1_Color_noise/internal/pkg/pin"
	"2020_1_Color_noise/internal/pkg/response"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	pinUsecase pin.IUsecase
}

func NewHandler(usecase pin.IUsecase) *Handler {
	return &Handler{
		pinUsecase: usecase,
	}
}

func (ph *Handler) Create(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	userId, ok := r.Context().Value("Id").(uint)
	if !ok {
		err := error.NoType.New("Received bad id from context")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	input := &models.InputPin{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		err = error.Wrap(err,"Decoding error during creation pin")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	_, err = govalidator.ValidateStruct(input)
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrapf(err, "request id: %s", "5"),
			"Name shouldn't be empty and longer 60 characters. " +
			"Description shouldn't be empty and longer 1000 characters. " +
			"Image should be base64")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	id, err := ph.pinUsecase.Create(input, userId)
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	resp := &models.ResponsePin{
		Id: id,
	}

	response.Respond(w, http.StatusCreated, resp)
}

func (ph *Handler) GetPin(w http.ResponseWriter, r *http.Request) {
	reqId:= r.Context().Value("ReqId")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting a pin"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	pin, err := ph.pinUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
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
	reqId:= r.Context().Value("ReqId")

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.WithMessage(error.BadRequest.Wrap(err, "Bad id in during getting pins for user"), "Bad id")
		error.ErrorHandler(w, error.Wrapf(err, "request id: %s", reqId))
		return
	}

	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pins, err := ph.pinUsecase.GetByUserId(uint(userId), start, limit)
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