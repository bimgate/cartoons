package main

var currentId int

var cartoons Cartoons
var episodes Episodes

// Give us some seed data
func init() {
	RepoCreateCartoon(Cartoon{Name: "Dilbert"})
	RepoCreateCartoon(Cartoon{Name: "xkcd"})
	RepoCreateCartoon(Cartoon{Name: "Snoopy"})

	RepoCreateCartoonEpisode(Episode{Name: "x", Episode_URL: "test_1.com"})
	RepoCreateCartoonEpisode(Episode{Name: "y", Episode_URL: "test_2.com"})
	RepoCreateCartoonEpisode(Episode{Name: "z", Episode_URL: "test_3.com"})
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
