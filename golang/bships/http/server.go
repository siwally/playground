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
const gamePath string = "/bships/games/"
const playerPath string = "/bships/games/players/"
const attackPath string = "/bships/games/attacks/"

// PlayerReq specifies the request body to POST when creating a player.
type PlayerReq struct {
	PlayerName string
	Ships      []bships.Ship
}

// AttackReq specifies the request body to POST when a player attackes another player.
type AttackReq struct {
	PlayerName string
	Move       bships.Coord
}

// Record of games in progress and players joined - expand for multi-game later.
var game *bships.Game

// StartServer creates an HTTP server listening for Battleships game set-up and move requests.
func StartServer(addr string) {
	log.Printf("Starting server at %v; handling:\n %v\n %v\n %v\n", addr, gamePath, playerPath, attackPath)

	http.HandleFunc(gamePath, GameHandler)
	http.HandleFunc(playerPath, PlayerHandler)
	http.HandleFunc(attackPath, AttackHandler)
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

	if err := game.AddPlayer(playerReq.PlayerName, playerReq.Ships...); err != nil {
		writeError(w, "Error adding player to game", err)
		return
	}

	writeJSON(w, fmt.Sprintf(`{"numPlayers": "%v"}`, len(game.Players)))
}

// AttackHandler processes RESTful requests to attack a player.
func AttackHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("AttackHandler invoked; path is %v", r.URL.Path)

	attackReq := AttackReq{}

	if err := json.NewDecoder(r.Body).Decode(&attackReq); err != nil {
		writeError(w, "Error decoding attack data", err)
		return
	}

	hit, sunk, err := game.Players[attackReq.PlayerName].Attack(attackReq.Move)

	if err != nil {
		writeError(w, "Error attacking player", err)
		return
	}

	writeJSON(w, fmt.Sprintf(`{"hit": "%v", "sunk": "%v"}`, hit, sunk != nil))
}

// Utility methods to avoid repetition

func writeJSON(w http.ResponseWriter, str string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, str)
}

func writeError(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(http.StatusInternalServerError)

	log.Printf(fmt.Sprintf("%s: %v", msg, err))
}
