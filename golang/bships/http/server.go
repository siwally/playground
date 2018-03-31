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

// PlayerReq specifies the request body to POST when creating a player.
type PlayerReq struct {
	PlayerName string
	Ships      []bships.Ship
}

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

	playerReq := PlayerReq{}
	err := json.NewDecoder(r.Body).Decode(&playerReq)

	log.Printf("Decoded PlayerReq data:`%v", playerReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error decoding game config: %v", err)
		return
	}

}

// TODO MoveHandler - can be underneath a player, and specify the player being attacked and the move as the payload
