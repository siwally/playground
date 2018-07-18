package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/siwally/golang/bships"
)

const player1Name string = "p1"
const player2Name string = "p2"

// TODO UI, so can play a game through and do some exploratory testing and think about evolution (before going too deep).
// TODO Tighten up, so need two players, have to alternate turns, can't attack yourself to cheat, etc.
// TODO Add light touch with identity, e.g. return tokens / ids for game and players.
// TODO Make tests data-driven at some point and set expectations in the table, e.g. hit, sunk, etc.

func TestGameplay(t *testing.T) {
	createGame(t)

	addPlayer(t, player1Name, bships.Ship{Start: bships.Coord{Row: 'B', Column: 1}, Dir: bships.TopToBottom, ShipType: mid})
	addPlayer(t, player2Name, bships.Ship{Start: bships.Coord{Row: 'C', Column: 2}, Dir: bships.LeftToRight, ShipType: dinky})

	attack(t, player1Name, bships.Coord{Row: 'A', Column: 1})
	attack(t, player1Name, bships.Coord{Row: 'B', Column: 1}) // hit
	attack(t, player1Name, bships.Coord{Row: 'B', Column: 2})
	attack(t, player1Name, bships.Coord{Row: 'C', Column: 1}) // hit
	attack(t, player1Name, bships.Coord{Row: 'D', Column: 1}) // hit
	attack(t, player1Name, bships.Coord{Row: 'E', Column: 1}) // hit, sunk
	attack(t, player1Name, bships.Coord{Row: 'E', Column: 1}) // hit, sunk
}

func DontTestServer(t *testing.T) {

	go StartServer(":8080")

	resp, err := http.Post("http://localhost:8080/bships/game", "application/json", nil)

	if err != nil {
		t.Errorf("Failed to post; err is %v", err)
	}

	resp.Write(os.Stdout)
}

// Small methods to simplify tests.

func createGame(t *testing.T) {
	req, _ := http.NewRequest("POST", "/bships/games", nil)

	resp := serve(req, GameHandler)

	checkResp(t, resp)
}

func addPlayer(t *testing.T, playerName string, ship bships.Ship) {
	player := PlayerReq{PlayerName: playerName, Ships: []bships.Ship{ship}}

	json, err := json.Marshal(player)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	req, _ := http.NewRequest("POST", "/bships/games/players", bytes.NewReader(json))
	resp := serve(req, PlayerHandler)

	checkResp(t, resp)
}

func attack(t *testing.T, playerName string, move bships.Coord) {
	attack := AttackReq{PlayerName: playerName, Move: move}

	json, err := json.Marshal(attack)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	req, _ := http.NewRequest("POST", "/bships/games/attacks", bytes.NewReader(json))

	resp := serve(req, AttackHandler)

	checkResp(t, resp)
}

func serve(req *http.Request, handlerFn func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFn)

	handler.ServeHTTP(rec, req)

	return rec
}

func checkResp(t *testing.T, resp *httptest.ResponseRecorder) {
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}

	// TODO Check expected outcome in body
	fmt.Printf("Response body: %v\n", resp.Body)
}
