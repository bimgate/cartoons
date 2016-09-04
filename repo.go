package main

import (
	"fmt"

	"github.com/gernest/nutz"
)

var currentId int

var cartoons Cartoons
var episodes Episodes

var dilbert_episode_name_print string

// Give us some seed data
func init() {
	/*
		/////////////////////////////////////vadi iz bazu
		databaseName := "my-database-dilbert-name.db"
		db := nutz.NewStorage(databaseName, 0600, nil)
		for i := 1; i < 10; i++ {

			key := (string(i))

			dilbert_episode_name := db.Get("dilbert", key)

			dilbert_episode_name_print := (string(dilbert_episode_name.Data))

			fmt.Println(dilbert_episode_name_print)
			///////////////////////////////////vadi iz bazu
	*/

	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Snoopy"})

	RepoCreateCartoonEpisode(Episode{Name: dilbert_episode_name_print, Episode_URL: ("http://cartoons-bimgate.rhcloud.com/static/" + dilbert_episode_name_print), Episode_share_URL: ("http://dilbert.com/strip/" + dilbert_episode_name_print)})
	//	RepoCreateCartoonEpisode(Episode{Name: "Episode_2", Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/2"})
	//	RepoCreateCartoonEpisode(Episode{Name: "Episode_3", Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/3"})
	//	RepoCreateCartoonEpisode(Episode{Name: "Episode_4", Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/4"})
	//	RepoCreateCartoonEpisode(Episode{Name: "Episode_5", Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/5"})
	//	RepoCreateCartoonEpisode(Episode{Name: "Episode_6", Episode_URL: "http://cartoons-bimgate.rhcloud.com/static/6"})

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

func RepoFindCartoonEpisode(id int) dilbert_episode_name_print {
	/////////////////////////////////////vadi iz bazu
	databaseName := "my-database-dilbert-name.db"
	db := nutz.NewStorage(databaseName, 0600, nil)
	for i := 1; i < 10; i++ {

		key := (string(i))

		dilbert_episode_name := db.Get("dilbert", key)

		dilbert_episode_name_print := (string(dilbert_episode_name.Data))

		fmt.Println(dilbert_episode_name_print)
		///////////////////////////////////vadi iz bazu
		return dilbert_episode_name_print
	}
	// return empty Episode if not found
	return dilbert_episode_name_print
}

func RepoCreateCartoon(c Cartoon) Cartoon {
	currentId += 1
	c.Id = currentId
	cartoons = append(cartoons, c)
	return c
}

func RepoCreateCartoonEpisode(e Episode) Episode {
	currentId += 1
	e.Id = currentId
	episodes = append(episodes, e)
	return e
}
