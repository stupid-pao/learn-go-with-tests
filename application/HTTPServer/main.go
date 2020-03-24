package main

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

type StubPlayStore struct {
	scores map[string]int
}

func (s *StubPlayStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

type InMemoryPlsyerStore struct{}

func (i *InMemoryPlsyerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlsyerStore{}}

	if err := http.ListenAndServe("localhost:5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
