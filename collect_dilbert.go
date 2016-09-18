package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

var next_comic_number_dilbert int
var dilbert_bolt_id int = 1
var v_string string
var u_d_format_to_string string

func collect_dilbert() {

	//Open DB

	db, err := bolt.Open("my-1-database-dilbert-name.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	/*
		//Firstly check if comic exist in Bolt db
		db.View(func(tx *bolt.Tx) error {
			//share_link_number_part := strconv.Itoa(i)
			b := tx.Bucket([]byte("dilbert"))
			c := b.Cursor()

			for k, v := c.Last(); k != nil; k, v = c.Next() {
				fmt.Printf("key=%s, value=%s\n", k, v)

				curent_last_comic_name_from_db := string(v)
				curent_last_comic_name_from_db_to_int, _ := strconv.Atoi(curent_last_comic_name_from_db)
				next_comic_name_dilbert = curent_last_comic_name_from_db_to_int + //NOVI NAREDNI DATUM a treba staviti i novi startni datum

			}

			return nil
		})
	*/
	//Firstly check if comic exist in Bolt db

	// set the starting date (in any way you wish)
	year, month, day := time.Now().Date() //Curent Day

	dayPicker := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	startDate := (dayPicker.Format("2006-01-02"))
	start, _ := time.Parse("2006-1-2", startDate)

	// handle error

	end, _ := time.Parse("2006-1-2", "2013-6-1")
	// handle error

	// set d to starting date and keep adding -1 day to it as long as Year doesn't change
	for d := start; d.Year() != end.Year(); d = d.AddDate(0, 0, -1) { // Day() or Month() or Year() for END
		// do stuff with d
		/*
			//Firstly check if comic exist in Bolt db
			db.View(func(tx *bolt.Tx) error {
				//share_link_number_part := strconv.Itoa(i)
				b := tx.Bucket([]byte("dilbert"))
				c := b.Cursor()

				for k, v := c.Last(); k != nil; {
					//fmt.Printf("key=%s, value=%s\n", k, v)
					u_d_format_to_string = (d.Format("2006-01-02"))

					v_string = string(v)

					fmt.Println("Datumi koji se menjaju" + u_d_format_to_string)
					fmt.Println("Poslednji datum ne menja se" + v_string)

				}

				return nil
			})



			//Firstly check if comic exist in Bolt db
		*/

		u := (d.Format("2006-01-02"))

		//fmt.Println(u)

		page_url := ("http://dilbert.com/strip/" + u)
		///////////Add to Bolt DB

		key := dilbert_bolt_id
		value := u

		dilb_bucket := []byte("dilbert")

		db.Update(func(tx *bolt.Tx) error {
			bucket, _ := tx.CreateBucketIfNotExists(dilb_bucket)

			err := bucket.Put([]byte(strconv.Itoa(key)), []byte(value))
			//fmt.Println("dbBolt Works  ", dilbert_bolt_id)
			return err

		})
		dilbert_bolt_id = dilbert_bolt_id + 1

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
			//fmt.Println("Success!")

		}

	}

}
