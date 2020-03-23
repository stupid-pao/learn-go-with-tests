package main

import (
	"fmt"
	"log"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, GetPlayerScore(player))

}

func GetPlayerScore(name string) string {
	switch name {
	case "pepper":
		return "20"
	case "Floyd":
		return "10"
	}
	return ""
}

func main() {
	handler := http.HandlerFunc(PlayerServer) // 注意这个声明  type HandlerFunc func(ResponseWriter, *Request)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
