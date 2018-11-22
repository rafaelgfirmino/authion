package router

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/oftall/authion/user/delivery"
	"github.com/oftall/authion/user/midleware"
	"net/http"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()
	router := Router.PathPrefix("/api/v1/security").Subrouter()
	addRoutes(router)
}

func addRoutes(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/signup").
		Handler(alice.New(midleware.ValidateUser, midleware.UserExist).Then(http.HandlerFunc(delivery.Signup)))
	router.Methods(http.MethodPost).Path("/signin").HandlerFunc(delivery.Signin)
	router.Methods(http.MethodPost).Path("/signup").HandlerFunc(delivery.Signup)
	router.Methods(http.MethodPut).Path("/signup/confirmation").HandlerFunc(delivery.ConfirmationToken)
}
