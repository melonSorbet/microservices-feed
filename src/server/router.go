package server

import (
	"feed/src/service"
	"net/http"
)

func RouteServer() {
	http.HandleFunc("GET /feed/friends", service.GetFriendsFeed)
	http.HandleFunc("GET /feed/explore", service.GetExplorePage)
}
