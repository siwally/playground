package main

import (
	"fmt"
	"net/http"
)

// GameHandler groups the HTTP handler functions
type GameHandler struct {
}

func main() {
	http.Handle("/bships/game", GameHandler{})
	http.ListenAndServe(":8080", nil)
}

func (GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received reqeust %v\n", r)

	// TOOD Add logging

	// TODO Take GameConfig as JSON input and use POST.  Create accompanying tests and can look in the logs at the payload.

	// TOOD Add player, which should include their name and ships - will involve changing core bships game.
}
