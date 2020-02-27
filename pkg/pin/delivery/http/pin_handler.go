package http

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"

	//"io"
	"net/http"
	"pinterest/pkg/models"
	"pinterest/pkg/pin/usecase"
	"strconv"
)

type Result struct {
	Status string      `json:"status"`
	Body interface{} `json:"body,omitempty"`
}

type PinHandler struct {
	pinUsecase usecase.IPinUsecase
}

func NewPinHandler(usecase usecase.IPinUsecase) *PinHandler {
	return &PinHandler{
		pinUsecase: usecase,
	}
}

func (ph *PinHandler) Add(w http.ResponseWriter, r *http.Request) {
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

	file, _, err  := r.FormFile("image")
	name := r.FormValue("name")
	description := r.FormValue("description")
	if err != nil || name == "" || description == "" {
		result.Status = "500"
		body["error"] = "Fill in all the fields"
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	image := bytes.NewBuffer(nil)
	_, err = io.Copy(image, file)
	if err != nil {
		result.Status = "500"
		body["error"] = "Internal error"
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

	var pin = &models.Pin{
		Name:    r.FormValue("name"),
		UserId:	 userId,
		Data: 	 image.Bytes(),
		Description: r.FormValue("description"),
	}

	id, err := ph.pinUsecase.Add(pin)
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Status = "200"
	body["id"] = id
	result.Body = body
	json.NewEncoder(w).Encode(result)
}

func (ph *PinHandler) GetPin(w http.ResponseWriter, r *http.Request) {
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

	pin, err := ph.pinUsecase.Get(uint(id))
	if err != nil {
		result.Status = "500"
		body["error"] = err.Error()
		result.Body = body
		json.NewEncoder(w).Encode(result)
		return
	}

	result.Status = "200"
	body["pin"] = pin
	result.Body = body
	json.NewEncoder(w).Encode(result)
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