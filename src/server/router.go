package server

import (
	"feed/src/service"
	"net/http"
)

func RouteServer(mux *http.ServeMux) {
	mux.HandleFunc("/feed/friends", service.GetFriendsFeed)
	mux.HandleFunc("/feed/explore", service.GetExplorePage)
}
