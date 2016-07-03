package main

import (
	"fmt"
	"io/ioutil"

	"net/http"
	"strconv"

	"io"

	"os"
	// "strings"
)

var y int = 1

func get_dilbert() {

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
}

func main() {
	get_dilbert()
}
