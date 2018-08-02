package delivery

import (
		"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request ){
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("olá mundo"))
	return
}