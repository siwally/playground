package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/siwally/golang/bships"
)

const gamePath string = "/bships/games"

const playerPath string = "/bships/games/players"

var game *bships.Game

// TODO Could have different types here that we expose, then marshall into game's types

// StartServer creates an HTTP server listening for Battleships game set-up and move requests.
func StartServer(addr string) {
	log.Printf("Starting server at %v; handling %v\n", addr, gamePath)

	http.HandleFunc(gamePath, GameHandler)
	http.HandleFunc(playerPath, PlayerHandler)
	http.ListenAndServe(addr, nil)
}

// GameHandler processes RESTful requests relating to the Battleships game set-up and state.
func GameHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Path is %v", r.URL.Path)

	cfg := bships.GameConfig{}
	err := json.NewDecoder(r.Body).Decode(&cfg)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error decoding game config: %v", err)
		return
	}

	game = bships.NewGame(cfg)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"gameId": "1"}`)
}

// PlayerHandler processes RESTful requests relating to player set-up.
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path is %v", r.URL.Path)
}

// TODO MoveHandler - can be underneath a player, and specify the player being attacked and the move as the payload
