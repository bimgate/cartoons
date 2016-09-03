package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

func get_dilbert() {

	// set the starting date (in any way you wish)
	year, month, day := time.Now().Date() //Curent Day

	dayPicker := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	startDate := (dayPicker.Format("2006-01-02"))
	start, _ := time.Parse("2006-1-2", startDate)

	end, _ := time.Parse("2006-1-2", "2015-6-1")
	// handle error

	// set d to starting date and keep adding 1 day to it as long as month doesn't change
	for d := start; (d.Year() != end.Year()) && (d.Month() != end.Month()); d = d.AddDate(0, 0, -1) {
		// do stuff with d

		u := (d.Format("2006-01-02"))

		page_url := ("http://dilbert.com/strip/" + u)

		tab_resp, _ := http.Get(page_url)
		tab_page, _ := ioutil.ReadAll(tab_resp.Body)

		tab_parsedPage, _ := gokogiri.ParseHtml(tab_page)

		img_tag := xpath.Compile("//*[@class='img-comic-link']/img")

		parsedPageSearch, _ := tab_parsedPage.Root().Search(img_tag)

		str := fmt.Sprint(parsedPageSearch)

		//fmt.Println(str) ////////////////////////////////

		var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)

		imgs := imgRE.FindAllStringSubmatch(str, -1)
		out := make([]string, len(imgs))
		//fmt.Println(out)

		for i := range out {
			out[i] = imgs[i][1]

			//fmt.Println(out[0])
			url_dilbert_cartoon := (out[0])

			response, e := http.Get(url_dilbert_cartoon)
			if e != nil {
				log.Fatal(e)
			}

			defer response.Body.Close()

			//open a file for writing
			/////////////////////////////////////////////////////////////

			file_path := "./static/dilbert/%v"
			f_path := fmt.Sprintf(file_path, u)

			file, err := os.Create(f_path)
			if err != nil {
				log.Fatal(err)
			}

			check := http.Client{
				CheckRedirect: func(r *http.Request, via []*http.Request) error {
					r.URL.Opaque = r.URL.Path
					return nil
				},
			}

			resp, err := check.Get(page_url) // add a filter to check redirect

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

			//FileServer

			//////////////////////////////////////////////////////////////////
			// Use io.Copy to just dump the response body to the file. This supports huge files
			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			fmt.Println("Success!")

		}

	}

}
