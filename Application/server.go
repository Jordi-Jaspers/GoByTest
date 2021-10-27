package main

import (
	"fmt"
	"net/http"
	"strings"
)

// The server.
type PlayerServer struct {
	store PlayerStore
}

// DB connection.
type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	PostPlayerScore(name string)
}

// Implementing the handler interface by implementing the serveHTTP method.
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.PostScore(w, player)

	case http.MethodGet:
		p.GetScore(w, player)

	}
}

// The GET call for the '/players/{name}' endpoint.
func (p *PlayerServer) GetScore(w http.ResponseWriter, player string) {
	score, ok := p.store.GetPlayerScore(player)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, score)
}

// The POST call for the '/players/{name}' endpoint.
func (p *PlayerServer) PostScore(w http.ResponseWriter, player string) {
	_, ok := p.store.GetPlayerScore(player)

	if !ok {
		p.store.GetPlayerScore(player)
	} else {
		// add User
	}

	w.WriteHeader(http.StatusAccepted)
}
