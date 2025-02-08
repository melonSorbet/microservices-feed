package server

import (
	"fmt"
	"log"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupServer() {
	mux := http.NewServeMux()
	RouteServer(mux)

	handler := CORS(mux)

	fmt.Println("Server running on port 8020")
	err := http.ListenAndServe(":8020", handler)
	if err != nil {
		log.Fatalf("Server is down: %v", err)
	}
}
