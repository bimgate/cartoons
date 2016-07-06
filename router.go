package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//////////////////////
	//fs := http.FileServer(http.Dir("./"))
	////////////////////////////////

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := http.Handler(http.FileServer(http.Dir("./")))

		handler := route.HandlerFunc
		handler := Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
