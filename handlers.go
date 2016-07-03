package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
	var y int = 1
	for i := 160000; i < 170000; i++ {

		rawURL := "http://cdn.ttgtmedia.com/rms/computerweekly/dt%v.png"
		url := fmt.Sprintf(rawURL, i)

		//fmt.Println(url)

		resp, _ := http.Get(url)

		page, _ := ioutil.ReadAll(resp.Body)

		if len(page) < 400 {
			goto next_number
		} else {
			fmt.Println(url) //for bolt

			//bolt

			file_name := strconv.Itoa(y)

			y++
			file, err := os.Create(file_name)

			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			defer file.Close()

			check := http.Client{
				CheckRedirect: func(r *http.Request, via []*http.Request) error {
					r.URL.Opaque = r.URL.Path
					return nil
				},
			}

			resp, err := check.Get(url) // add a filter to check redirect

			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			defer resp.Body.Close()
			fmt.Println(resp.Status)

			size, err := io.Copy(file, resp.Body)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s with %v bytes downloaded" /*fileName,*/, size)

			//bolt

		}
	next_number:
	}

	//

}

func FileServer(w http.ResponseWriter, r *http.Request) {
	http.Dir("./")
	http.ListenAndServe("OPENSHIFT_GO_PORT", nil)
}

func CartoonEpisodeShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	episodeId := vars["episodeId"]
	fmt.Fprintln(w, "Episode show:", episodeId)
}
