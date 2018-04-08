package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/siwally/golang/bships"
)

// RESTful paths exposed by this service
const gamePath string = "/bships/games"
const playerPath string = "/bships/games/players"

// PlayerReq specifies the request body to POST when creating a player.
type PlayerReq struct {
	PlayerName string
	Ships      []bships.Ship
}

// Record of games in progress and players joined - expand for multi-game
var game *bships.Game
var numPlayers int

// StartServer creates an HTTP server listening for Battleships game set-up and move requests.
func StartServer(addr string) {
	log.Printf("Starting server at %v; handling %v\n", addr, gamePath)

	http.HandleFunc(gamePath, GameHandler)
	http.HandleFunc(playerPath, PlayerHandler)
	http.ListenAndServe(addr, nil)
}

// GameHandler processes RESTful requests relating to the Battleships game set-up and state.
func GameHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GameHandler invoked; path is %v", r.URL.Path)

	game = bships.NewGame(cfg)

	writeJSON(w, `{"gameId": "1"}`)
}

// PlayerHandler processes RESTful requests relating to player set-up.
func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("PlayerHandler invoked; path is %v", r.URL.Path)

	playerReq := PlayerReq{}

	if err := json.NewDecoder(r.Body).Decode(&playerReq); err != nil {
		writeError(w, "Error decoding player data", err)
		return
	}

	log.Printf("Decoded PlayerReq data:`%v", playerReq)

	if err := game.AddPlayer(playerReq.PlayerName, playerReq.Ships...); err != nil {
		writeError(w, "Error adding player to game", err)
		return
	}

	numPlayers++
	writeJSON(w, fmt.Sprintf(`{"playerId": "%v"}`, numPlayers))
}

// TODO MoveHandler - can be underneath a player, and specify the player being attacked and the move as the payload

// Utility methods to avoid repetition

func writeJSON(w http.ResponseWriter, str string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, str)
}

func writeError(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	log.Fatalf(fmt.Sprintf("%s: %v", msg, err))
}
