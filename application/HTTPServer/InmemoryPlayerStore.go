package main

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func NewInMemoryPlayerStore() *InMemoryPlsyerStore {
	return &InMemoryPlsyerStore{map[string]int{}}
}

type InMemoryPlsyerStore struct {
	store map[string]int
}

func (i *InMemoryPlsyerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlsyerStore) RecordWin(name string) {
	i.store[name]++
}
