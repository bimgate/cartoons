package main

import (
	"fmt"

	"github.com/gernest/nutz"
)

//var currentId int

var cartoons Cartoons
var episodes Episodes

// Give us some seed data
func init() {

	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Dilbert", Number_of_Episodes: 120, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/dilbert/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "xkcd", Number_of_Episodes: 100, Episodes_URL: "http://cartoons-bimgate.rhcloud.com/static/xkcd/"})
	RepoCreateCartoon(Cartoon{Bootstrap_URL: "http://cartoons-bimgate.rhcloud.com", Name: "Snoopy"})

	RepoCreateCartoonEpisode(Episode{Name: dilbert_episode_name_print, Episode_URL: ("http://cartoons-bimgate.rhcloud.com/static/" + dilbert_episode_name_print), Episode_share_URL: ("http://dilbert.com/strip/" + dilbert_episode_name_print)})

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

func RepoFindCartoonEpisode(id int) Episode {

	for _, e := range episodes {
		if e.Id == id {
			return e
		}
	}

	// return empty Episode if not found
	return Episode{}
}

/*
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
*/
