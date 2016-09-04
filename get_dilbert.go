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
	// handle error

	end, _ := time.Parse("2006-1-2", "2015-6-1")
	// handle error

	// set d to starting date and keep adding -1 day to it as long as Year doesn't change
	for d := start; d.Month() != end.Month(); d = d.AddDate(0, 0, -1) { //Month() or Year()
		// do stuff with d
		dilbert_bolt_id := 1
		u := (d.Format("2006-01-02"))

		fmt.Print(u)

		page_url := ("http://dilbert.com/strip/" + u)
		///////////Add to Bolt DB
		//Open DB
		db, err := bolt.Open("my-database-dilbert-name.db", 0600, nil)

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		key := dilbert_bolt_id
		value := u

		dilb_bucket := []byte("dilbert")

		err = db.Update(func(tx *bolt.Tx) error {
			bucket, _ := tx.CreateBucketIfNotExists(dilb_bucket)

			err := bucket.Put([]byte(strconv.Itoa(key)), []byte(strconv.Itoa(value)))
			fmt.Print("dbBolt Works  ", dilbert_bolt_id)
			return err

		})

		///////////Add to Bolt DB
		tab_resp, _ := http.Get(page_url)
		tab_page, _ := ioutil.ReadAll(tab_resp.Body)

		tab_parsedPage, _ := gokogiri.ParseHtml(tab_page)

		img_tag := xpath.Compile("//*[@class='img-comic-link']/img")

		parsedPageSearch, _ := tab_parsedPage.Root().Search(img_tag)

		str := fmt.Sprint(parsedPageSearch)

		var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)

		imgs := imgRE.FindAllStringSubmatch(str, -1)
		out := make([]string, len(imgs))

		for i := range out {
			out[i] = imgs[i][1]

			url_dilbert_cartoon := (out[0])

			response, e := http.Get(url_dilbert_cartoon)
			if e != nil {
				log.Fatal(e)
			}

			defer response.Body.Close()

			//open a file for writing
			file_path := "./static/dilbert/%v"
			f_path := fmt.Sprintf(file_path, u)

			file, err := os.Create(f_path)
			if err != nil {
				log.Fatal(err)
			}

			//FileServer

			_, err = io.Copy(file, response.Body)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			fmt.Println("Success!")

		}

	}

}
