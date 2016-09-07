package main

import (
	"fmt"

	//"log"
	//"strconv"

	"github.com/boltdb/bolt"

	//"github.com/gernest/nutz"
)

var currentId int
var cartoons Cartoons
var episodes Episodes

var databaseName = "my-database-dilbert-name.db"
var db, _ = bolt.Open(databaseName, 0600, nil)

func init() {
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})

}

func init() {

	/////////////////////////////////////vadi iz bazu
	//Open DB
	for i := 1; i < 10; i++ {

		// retrieve the data
		_ = db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(dilbertBucket)
			if bucket == nil {
				return fmt.Errorf("Bucket %q not found!", dilbertBucket)
			}

			val := bucket.Get([]byte("1"))
			fmt.Println(string(val))

			RepoCreateCartoonEpisode(Episode{Name: "vtest", Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/XZY", Episode_share_URL: "SHARE_URL"})

			return nil
		})

		//	RepoCreateCartoonEpisode(Episode{Name: m, Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/XZY", Episode_share_URL: "SHARE_URL"})

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
