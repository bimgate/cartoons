package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

func get_xkcd() {

	var y int = 1
	for i := 100; i < 200; i++ {
		///////////////////////////////////////////////////////////////////////////////////

		url := "http://xkcd.com/%v/"

		page_url := fmt.Sprintf(url, i)

		tab_resp, _ := http.Get(page_url)
		tab_page, _ := ioutil.ReadAll(tab_resp.Body)

		tab_parsedPage, _ := gokogiri.ParseHtml(tab_page)
		img_tag := xpath.Compile("//*[@id='comic']/img")

		parsedPageSearch, _ := tab_parsedPage.Root().Search(img_tag)

		str := fmt.Sprint(parsedPageSearch)

		var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)

		imgs := imgRE.FindAllStringSubmatch(str, -1)
		out := make([]string, len(imgs))

		for i := range out {
			out[i] = imgs[i][1]

			//fmt.Println("http:"+out[0])
			url_xkcd_cartoon := ("http:" + out[0])

			response, e := http.Get(url_xkcd_cartoon)
			if e != nil {
				log.Fatal(e)
			}

			defer response.Body.Close()

			//open a file for writing

			file_path := "./static/xkcd/%v"
			f_path := fmt.Sprintf(file_path, y)

			y++

			file, err := os.Create(f_path)
			if err != nil {
				log.Fatal(err)
			}
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
