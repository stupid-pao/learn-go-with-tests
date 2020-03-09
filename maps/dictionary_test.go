package maps

import (
	"errors"
	"testing"
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

//var ErrNotFound = DictionaryErr.Error("could not find the word you were looking for")
//var ErrWordExists = DictionaryErr.Error("cannot add word because it already exists")
//var ErrWordNotExists = DictionaryErr.Error("can not update word because it does not exists")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// map 又一个有趣的特性，不实用指针就可以修改他们。这是因为 map 是饮用类型。
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want %s", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}

}

/*
map 作为引用类型很好，因为无论map有多大，都只会有一个副本。
但是尝试使用一个nil的map会得到指针异常，
所以永远不要初始化一个空的 map 像这样： var m map[string]string
正确使用方式：
	dictionary := map[string]string{}
	or
	dictionary := make[map[string]string]
*/
