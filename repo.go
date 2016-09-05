package main

import (
	"fmt"

	"github.com/boltdb/bolt"
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

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("2"))
		fmt.Printf("The answer is: %s\n", v)
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
