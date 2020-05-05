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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Real-Status", fmt.Sprint(status))
	response := &Response{
		Status: status,
		Body:   body,
	}

	easyjson.MarshalToHTTPResponseWriter(response, w)
	return
}
