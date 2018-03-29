package http

import (
	"fmt"
	"log"
	"net/http"
)

const gamePath string = "/bships/game"

// GameHandler groups the HTTP handler functions
type GameHandler struct {
}

// StartServer creates an HTTP server listening for Battleships game set-up and move requests.
func StartServer(addr string) {
	log.Printf("Starting server at %v; handling %v\n", addr, gamePath)

	http.Handle(gamePath, GameHandler{})
	http.ListenAndServe(addr, nil)
}

func (GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling reqeust %v\n", r)

	fmt.Printf("Handling reqeust %v\n", r)

	// TODO Take GameConfig as JSON input and use POST.  Create accompanying tests and can look in the logs at the payload.

	// TOOD Add player, which should include their name and ships - will involve changing core bships game.

	log.Println("Finished handling reqeust")
}
