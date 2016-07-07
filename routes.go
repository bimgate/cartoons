package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CartoonIndex",
		"GET",
		"/cartoons",
		CartoonIndex,
	},
	Route{
		"DownloadCartoon",
		"GET",
		"/cartoons/static/",
		DownloadCartoon,
	},
	Route{
		"CartoonEpisodesIndex",
		"GET",
		"/cartoons/{cartoonId}/episodes",
		CartoonEpisodesIndex,
	},
	Route{
		"CartoonEpisodeShow",
		"GET",
		"/cartoons/{cartoonId}/episodes/{episodeId}",
		CartoonEpisodeShow,
	},
	Route{
		"static",
		"GET",
		"/static",
		Static,
	},
}
