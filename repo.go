package main

import (
	"fmt"

	//"log"
	//"strconv"

	//"github.com/boltdb/bolt"

	"github.com/gernest/nutz"
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

	RepoCreateCartoonEpisode(Episode{Name: val_print, Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/XZY", Episode_share_URL: "SHARE_URL"})

}

func RepoCreateCartoonEpisode(e Episode) Episode {

	/////////////////////////////////////vadi iz bazu
	//Open DB
	for i := 1; i < 10; i++ {
		databaseName := "my-1-database-dilbert-name.db"
		db := nutz.NewStorage(databaseName, 0600, nil)

		n := db.Get("dilbert", "11")

		m := (string(n.Data))
		val_print = fmt.Sprintf("%s", m)
	}

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
