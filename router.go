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
		var handler http.Handler
		fs := http.FileServer(route)
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = fs
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
