package main

import (
	"fmt"

	"log"

	"github.com/boltdb/bolt"
)

var currentId int
var cartoons Cartoons
var episodes Episodes

func init() {
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})

}

func init() {

	/////////////////////////////////////vadi iz bazu
	//Open DB
	db, err := bolt.Open("my-database-dilbert-name.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("MyBucket"))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)

			RepoCreateCartoonEpisode(Episode{Name: "test-1", Episode_URL: ("http://cartoons-bimgate.rhcloud.com/static/"), Episode_share_URL: ("http://dilbert.com/strip/")})

			return nil
		})
		return nil
	})

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
