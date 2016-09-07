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

var databaseName = "my-database-dilbert-name.db"

func init() {
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})

}

func init() {

	/////////////////////////////////////vadi iz bazu
	//Open DB
	for i := 1; i < 10; i++ {

		db, err := bolt.Open(databaseName, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		i_to_string := strconv.Itoa(i)
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("dilbert"))
			v := b.Get([]byte(i_to_string))
			//fmt.Printf("The answer is: %s\n", v)

			m := string(v)

			fmt.Println(m)

			RepoCreateCartoonEpisode(Episode{Name: m, Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/xzy", Episode_share_URL: "SHARE_URL"})
			return nil
		})

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
