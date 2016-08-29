package main

type Cartoon struct {
	Id                 int    `json: "id"`
	Name               string `json: "name"`
	Number_of_Episodes int    `json: "episodes_number"`
	Episodes_URL       string `json: "episode_url"`
}

type Cartoons []Cartoon
