package http

import (
	"pinterest/internal/pkg/error"
	"pinterest/internal/pkg/pin"
	"pinterest/internal/pkg/response"
	"strconv"

	"encoding/json"
	"github.com/gorilla/mux"

	"net/http"
	"pinterest/internal/models"
)

type Result struct {
	Status string      `json:"status"`
	Body interface{} `json:"body,omitempty"`
}

type Handler struct {
	pinUsecase pin.IUsecase
}

func NewHandler(usecase pin.IUsecase) *Handler {
	return &Handler{
		pinUsecase: usecase,
	}
}

func (ph *Handler) Add(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, err)
		return
	}

	input := &models.InputPin{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		error.ErrorHandler(w, err)
		return
	}

	id, err := ph.pinUsecase.Add(input)
	if err != nil {
		error.ErrorHandler(w, err)
		return
	}

	resp := &models.ResponsePin{
		Id: id,
	}

	response.Respond(w, http.StatusCreated, resp)
}

func (ph *Handler) GetPin(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value("isAuth") == false {
		err := error.Unauthorized.New("User is unauthorized")
		error.ErrorHandler(w, err)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		err = error.BadRequest.Wrap(err, "Bad id when getting a pin")
		error.ErrorHandler(w, err)
		return
	}

	pin, err := ph.pinUsecase.GetById(uint(id))
	if err != nil {
		error.ErrorHandler(w, err)
		return
	}

	resp := &models.ResponsePin{
		Id:          pin.Id,
		UserId:      pin.UserId,
		Name:        pin.Name,
		Description: pin.Description,
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