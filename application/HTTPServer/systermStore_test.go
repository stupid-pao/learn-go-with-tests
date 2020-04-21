package main

import (
	"io"
	"strings"
	"testing"
)

type FileSystemStore struct {
	database io.ReadSeeker
}

func (f *FileSystemStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemStore) GetPlayerStore(name string) int {

	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		// NewReader 也实现了 ReadSeeker，所以可以用
		database := strings.NewReader(`[
		    {"Name": "Cleo", "Wins": 10},
		    {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
		    {"Name": "Cleo", "Wins": 10},
		    {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemStore{database}

		got := store.GetPlayerStore("Chris")

		want := 33

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
