package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rafaelgfirmino/authion/user/delivery"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()
	router := Router.PathPrefix("/api/v1/security").Subrouter()
	addRoutes(router)
}

func addRoutes(router *mux.Router){
	router.Methods(http.MethodGet).Path("/signup").HandlerFunc(delivery.Signup)
}