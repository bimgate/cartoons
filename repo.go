package main

import (
	"fmt"

	"github.com/gernest/nutz"
)

var currentId int
var cartoons Cartoons
var episodes Episodes

func init() {
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})

}

func RepoFindCartoon(id int) Cartoon {
	for _, c := range cartoons {
		if c.Id == id {
			return c
		}
	}
	// return empty Cartoon if not found
	return Cartoon{}
}

func RepoCreateCartoon(c Cartoon) Cartoon {
	currentId += 1
	c.Id = currentId
	cartoons = append(cartoons, c)
	return c
}

func RepoFindCartoonEpisode(id int) Episode {

	/////////////////////////////////////vadi iz bazu
	databaseName := "my-database-dilbert-name.db"
	db := nutz.NewStorage(databaseName, 0600, nil)
	for i, e := range episodes {
		if e.Id == id {

			key := (string(i))

			dilbert_episode_name := db.Get("dilbert", key)

			dilbert_episode_name_print := (string(dilbert_episode_name.Data))

			fmt.Println(dilbert_episode_name_print)
			///////////////////////////////////vadi iz bazu
			return e
		}

	}
	// return empty Episode if not found
	return Episode{}
}
