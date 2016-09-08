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
	err := http.ListenAndServe(bind, router)

	//err := http.ListenAndServe(":9090", router)

	if err != nil {
		panic(err)
	}
}
