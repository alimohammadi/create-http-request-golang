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

func main() {
	api := &api{addr: ":8080"}

	srv := &http.Server{Addr: api.addr, Handler: api}

	srv.ListenAndServe()
}
