package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	/*
		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
		http.Handle("/", router)
	*/

	static := http.FileServer(http.Dir("static"))
	ruter.Handle("/static/", http.StripPrefix("/static/", static))

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
