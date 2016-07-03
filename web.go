package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	router := NewRouter()

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, http.FileServer(http.Dir("/")))

	if err != nil {
		panic(err)
	}
}
