package main

import (
	"fmt"

	"log"
	"strconv"

	"github.com/boltdb/bolt"

	//"github.com/gernest/nutz"
)

var currentId int
var cartoons Cartoons
var episodes Episodes

var val_print string

func init() {
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})

}

func init() {

	/////////////////////////////////////vadi iz bazu
	//Open DB

	db, err := bolt.Open("my-2-database-dilbert-name.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for i := 1; i < 10; i++ {

		// retrieve the data
		db.View(func(tx *bolt.Tx) error {

			dil_bu := []byte("dilbert-2")

			b, _ := tx.CreateBucketIfNotExists(dil_bu)

			v_string := strconv.Itoa(i)

			v := b.Get([]byte(v_string))

			val_print = fmt.Sprintf("%s", v)

			RepoCreateCartoonEpisode(Episode{Name: val_print, Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/XZY", Episode_share_URL: "SHARE_URL"})

			return nil
		})

		//RepoCreateCartoonEpisode(Episode{Name: val_print, Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/XZY", Episode_share_URL: "SHARE_URL"})

	}
}

func RepoCreateCartoonEpisode(e Episode) Episode {
	currentId += 1
	e.Id = currentId
	episodes = append(episodes, e)
	return e
}

func RepoCreateCartoon(c Cartoon) Cartoon {
	currentId += 1
	c.Id = currentId
	cartoons = append(cartoons, c)
	return c
}
