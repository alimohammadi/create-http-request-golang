package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server\n"))

	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			w.Write([]byte("hello from the main route \n"))
			return
		case "/users":
			w.Write([]byte("hello from the users route \n"))
			return
		default:
			w.Write([]byte("404 page"))
			return
		}

	default:
		w.Write([]byte("404 page"))
		return
	}
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users list"))
}

func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created user"))
}

func main() {
	api := &api{addr: ":8080"}

	// Initialize the server mux
	mux := http.NewServeMux()
	srv := &http.Server{Addr: api.addr, Handler: api}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	srv.ListenAndServe()
}
