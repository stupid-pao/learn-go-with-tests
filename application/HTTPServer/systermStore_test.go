package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

type FileSystemStore struct {
	database *json.Encoder
	league   League
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem form getting info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {

	err := initialisePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initiallising player db file, %v", err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem from loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{
		database: json.NewEncoder(&Tape{file}),
		league:   league,
	}, nil

}

func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}

	f.database.Encode(f.league)

}

func TestFileSystemStore(t *testing.T) {

	t.Run("/league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		    {"Name": "Cleo", "Wins": 10},
		    {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := NewFileSystemStore(database)

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
		database, cleanDatabase := createTempFile(t, `[
		    {"Name": "Cleo", "Wins": 10},
		    {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := NewFileSystemStore(database)

		got := store.GetPlayerScore("Chris")

		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		    {"Name": "Cleo", "Wins": 10},
		    {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := NewFileSystemStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")

		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		    {"Name": "Cleo", "Wins": 10},
		    {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := NewFileSystemStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t *testing.T, got, want int) {

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile

}
