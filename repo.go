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

	/////////////////////////////////////vadi iz bazu
	databaseName := "my-database-dilbert-name.db"
	db := nutz.NewStorage(databaseName, 0600, nil)
	for i := 1; i < 50; i++ {

		key := (string(i))

		dilbert_episode_name := db.Get("dilbert", key)

		dilbert_episode_name_print := (string(dilbert_episode_name.Data))

		fmt.Println(dilbert_episode_name_print)
		episodes = dilbert_episode_name_print
		//////////////////////////////////////////////////
		//Episode{Name: dilbert_episode_name_print, Episode_URL: ("http://cartoons-bimgate.rhcloud.com/static/" + dilbert_episode_name_print), Episode_share_URL: ("http://dilbert.com/strip/" + dilbert_episode_name_print)}

		///////////////////////////////////vadi iz bazu

	}
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
