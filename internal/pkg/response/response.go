package response

import (
	"2020_1_Color_noise/internal/pkg/metric"
	"fmt"
	"github.com/mailru/easyjson"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body,omitempty"`
}

func Respond(w http.ResponseWriter, status int, body interface{}, path string) {
	w.Header().Set("Content-Type", "application/json")
	response := &Response{
		Status: status,
		Body:   body,
	}
	metric.IncreaseRps(fmt.Sprint(status), path)
	easyjson.MarshalToHTTPResponseWriter(response, w)
	return
}
