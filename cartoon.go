package main

type Cartoon struct {
	Id                  int    `json: "id"`
	Name                string `json: "name"`
	Number_of_Episodes  int    `json: "episodes_number"`
	Dilbert_Episode_URL string `json: "episode_url"`
	Xkcd_Episode_URL    string `json: "episode_url"`
}

type Cartoons []Cartoon
