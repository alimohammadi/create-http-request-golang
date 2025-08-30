package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	// Initialize the server mux
	mux := http.NewServeMux()
	srv := &http.Server{Addr: api.addr, Handler: mux}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
