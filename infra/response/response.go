package response

import (
	"encoding/json"
	"net/http"
)
type data struct {
	Data interface{}
}
func Json(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}