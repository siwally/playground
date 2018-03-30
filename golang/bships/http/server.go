package http

import (
	"io"
	"log"
	"net/http"
)

const gamePath string = "/bships/game"

// StartServer creates an HTTP server listening for Battleships game set-up and move requests.
func StartServer(addr string) {
	log.Printf("Starting server at %v; handling %v\n", addr, gamePath)

	http.HandleFunc(gamePath, GameHandler)
	http.ListenAndServe(addr, nil)
}

func GameHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling requust %v\n", r)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// TODO Take GameConfig as JSON input and use POST.  Create accompanying tests and can look in the logs at the payload.

	// TOOD Add player, which should include their name and ships - will involve changing core bships game.

	io.WriteString(w, `{"key1": "value1"}`)

	log.Printf("Finished handling request %v\n", r)
}
