package cmd

import (
	"fmt"
	"net/http"

	"github.com/aditya-nagare/take-home-test/problem-2/handlers"
	"github.com/aditya-nagare/take-home-test/problem-2/middleware"
	"github.com/aditya-nagare/take-home-test/problem-2/routes"
	"github.com/gorilla/mux"
)

func Run() {
	url := "http://reqres.in/api/users/"
	handlers := handlers.NewHandlerImplementation(url)
	router := mux.NewRouter()
	router.Use(middleware.Middleware)
	routes.NewRoutes(router, handlers)

	fmt.Println("Server Running on port :9001")
	http.ListenAndServe(":9001", router)
}
