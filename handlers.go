package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Cartoones Kingdom!")

}

func CartoonIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cartoons); err != nil {
		panic(err)
	}

}

func CartoonEpisodesIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(episodes); err != nil {
		panic(err)
	}

}

func DownloadCartoon(w http.ResponseWriter, r *http.Request) {
	//
	get_dilbert()
	//

}

func CartoonEpisodeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	fmt.Fprintln(w, "Episode show:", episodeId)
}

func Static(w http.ResponseWriter, r *http.Request) {

}
