package response

import (
	"fmt"
	"github.com/mailru/easyjson"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}

func Respond(w http.ResponseWriter, status int, body interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error while marshalling, error: ", r)
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Real-Status", fmt.Sprint(status))
	response := &Response{
		Status: status,
		Body:   body,
	}

	_, _, err := easyjson.MarshalToHTTPResponseWriter(response, w)
	if err != nil {
		panic(err)
	}
	return
}
