package main

type Cartoon struct {
	Id                 int    `json: "id"`
	Name               string `json: "name"`
	Number_of_Episodes int    `json: "episodes_number"`
	Episode_id         int    `json: "episode_id"`
}

type Cartoons []Cartoon
