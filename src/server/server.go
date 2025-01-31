package server

import (
	"fmt"
	"net/http"
)

func SetupServer() {
	// TODO:
	RouteServer()
	fmt.Println("Server running on 8020")
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		println("server is down")
		return
	}
}
