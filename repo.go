package main

import (
	"fmt"

	"github.com/gernest/nutz"
)

var cartoons Cartoons
var episodes Episodes

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
			return dilbert_episode_name_print
		}

	}
	// return empty Episode if not found
	return Episode{}
}
