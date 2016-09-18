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
	"strings"

	"github.com/boltdb/bolt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
)

var c_value string
var next_comic_number_xkcd int

func get_xkcd() {
	//Open DB

	db, err := bolt.Open("my-1-database-dilbert-name.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//At first scrip Start part - Firstly check if comic exist in Bolt db - should to be commented and i in fore loop should have start Number values

	//Firstly check if comic exist in Bolt db
	db.View(func(tx *bolt.Tx) error {
		//share_link_number_part := strconv.Itoa(i)
		b := tx.Bucket([]byte("xkcd"))
		c := b.Cursor()

		for k, v := c.Last(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)

			curent_last_comic_number_from_db := string(k)
			curent_last_comic_number_from_db_to_int, _ := strconv.Atoi(curent_last_comic_number_from_db)
			next_comic_number_xkcd = curent_last_comic_number_from_db_to_int + 1

		}

		return nil
	})

	//Firstly check if comic exist in Bolt db

	for i := (next_comic_number_xkcd - 1); i <= next_comic_number_xkcd; i++ {

		url := "http://xkcd.com/%v/"

		page_url := fmt.Sprintf(url, i)

		tab_resp, _ := http.Get(page_url)
		tab_page, _ := ioutil.ReadAll(tab_resp.Body)

		tab_parsedPage, _ := gokogiri.ParseHtml(tab_page)
		///////////////////////////////////////////////////////////////////xkcd_episode_name
		collect_xkcd_episode_name := xpath.Compile(".//*[@id='ctitle']/text()")
		parsedPageSearch_xkcd_episode_name, _ := tab_parsedPage.Root().Search(collect_xkcd_episode_name)
		xkcd_episode_name := fmt.Sprint(parsedPageSearch_xkcd_episode_name)

		xkcd_episode_name_for_db := strings.TrimLeft(strings.TrimRight(xkcd_episode_name, "]"), "[")
		fmt.Println(xkcd_episode_name_for_db)

		///////////Add to Bolt DB
		xkcd_id := i
		key := xkcd_id

		c_value := xkcd_episode_name_for_db
		value := c_value //vrednost neka bude ime

		xkcd_bucket := []byte("xkcd")

		db.Update(func(tx *bolt.Tx) error {
			bucket, _ := tx.CreateBucketIfNotExists(xkcd_bucket)

			err := bucket.Put([]byte(strconv.Itoa(key)), []byte(value))
			fmt.Println("dbBolt Works  ", xkcd_id)
			return err

		})

		///////////Add to Bolt DB

		////////////////////////////////////////////////////////////////imgs
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
			f_path := fmt.Sprintf(file_path, c_value)

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
