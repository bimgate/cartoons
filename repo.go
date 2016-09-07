package main

import (
	"fmt"

	//"github.com/boltdb/bolt"
	"strconv"

	"github.com/gernest/nutz"
)

var currentId int
var cartoons Cartoons
var episodes Episodes

var databaseName = "my-database-dilbert-name.db"
var db = nutz.NewStorage(databaseName, 0600, nil)

func init() {
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})

}

func init() {

	/////////////////////////////////////vadi iz bazu
	//Open DB
	for i := 1; i < 10; i++ {

		db := nutz.Open(databaseName, 0600, nil)

		i_to_string := strconv.Itoa(i)
		n := db.Get("dilbert", i_to_string)

		m := (string(n.Data))

		fmt.Println(m)

		RepoCreateCartoonEpisode(Episode{Name: m, Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/", Episode_share_URL: "SHARE_URL"})
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
