package main

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &Tape{file}
	tape.Write([]byte("abc"))

	file.Seek(0, 0) // cursor 归零，之前进行过写入操作从cursor 的位置开始读，会读不到写入的数据
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}
