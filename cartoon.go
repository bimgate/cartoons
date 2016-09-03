package main

type Cartoon struct {
	Bootstrap_URL      string `json: "bootstrap_url"`
	Id                 int    `json: "id"`
	Name               string `json: "name"`
	Number_of_Episodes int    `json: "episodes_number"`
	Episodes_URL       string `json: "episode_url"`
	Bootstrap_URL      string `json: "bootstrap_url"`
}

type Cartoons []Cartoon
