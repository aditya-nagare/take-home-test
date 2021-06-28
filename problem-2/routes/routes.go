package routes

import (
	"github.com/aditya-nagare/take-home-test/problem-2/handlers"
	"github.com/gorilla/mux"
)

//NewRoutes returns list of routes
func NewRoutes(router *mux.Router, handler *handlers.HandlerImplementation) {
	router.HandleFunc("/user/{ID}", handler.GetUserEmail).Methods("GET")
}
