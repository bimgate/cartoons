package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//////////////////////
	//http.Handler(http.FileServer(http.Dir("./")))
	//fs := http.FileServer(http.Dir("./"))
	////////////////////////////////

	router := mux.NewRouter().StrictSlash(true)
	//router.Handle("/", http.FileServer(http.Dir("./")))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
