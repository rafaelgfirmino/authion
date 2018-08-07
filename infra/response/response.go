package response

import (
	"encoding/json"
	"net/http"
)

func Json(response interface{}, w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	return
}
