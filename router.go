package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	//////////////////////
	fs := http.FileServer(http.Dir("./"))
	////////////////////////////////

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handlerfs http.Handler

		handler := handlerfs(fs)

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		//handler = FileServer(handler, http.FileServer(http.Dir("./")))
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
