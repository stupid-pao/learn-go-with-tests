package main

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// 有了 FileSystemStore 就不需要这个了
// func NewInMemoryPlayerStore() *InMemoryPlayerStore {
// 	return &InMemoryPlayerStore{map[string]int{}}
// }

// type InMemoryPlayerStore struct {
// 	store map[string]int
// }

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
// 	return i.store[name]
// }

// func (i *InMemoryPlayerStore) RecordWin(name string) {
// 	i.store[name]++
// }

// func (i *InMemoryPlayerStore) GetLeague() League {
// 	var league []Player
// 	for name, wins := range i.store {
// 		league = append(league, Player{name, wins})
// 	}
// 	return league
// }
