package main

type Episode struct {
	Id                int    `json: "id"`
	Name              string `json: "name"`
	Episode_URL       string `json: "episode_url"`
	Episode_share_URL string `json: "episode_url"`
}

type Episodes []Episode
