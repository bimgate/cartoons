package main

type Cartoon struct {
	Id                 int    `json: "id"`
	Name               string `json: "name"`
	Number_of_Episodes int    `json: "episodes_number"`
}

type Cartoons []Cartoon
