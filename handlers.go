package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//var cartoonBucket string

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

func CartoonDilbertEpisodesIndex(w http.ResponseWriter, r *http.Request) {
	currentId = 0
	cartoonBucket_dilbert := "dilbert"
	GetFromBoltDb(cartoonBucket_dilbert)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(episodes); err != nil {
		panic(err)

	}
	episodes = nil
}

func CartoonXkcdEpisodesIndex(w http.ResponseWriter, r *http.Request) {
	currentId = 0
	cartoonBucket_xkcd := "xkcd"
	GetFromBoltDb(cartoonBucket_xkcd)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(episodes); err != nil {
		panic(err)

	}
	episodes = nil
}

func DownloadCartoon(w http.ResponseWriter, r *http.Request) {
	//
	collect_dilbert()

	//

}

func DownloadXkcd(w http.ResponseWriter, r *http.Request) {
	//
	get_xkcd()

	//

}

func CartoonEpisodeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	fmt.Fprintln(w, "Episode show:", episodeId)
}

func StaticDilbertXkcd(w http.ResponseWriter, r *http.Request) {

}
