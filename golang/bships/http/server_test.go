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

// TODO Build up test to adding and attacking multiple players, and make data-driven to play out a whole game.
// TODO Then tighten up, so need two players, have to alternate turns, can't attack yourself to cheat, etc.
// TODO How far do we want to go with identity though?  Interesting to see how transparent we could make it.
// (Presumably playerId could be a token that's returned from adding a player, or earlier on, then it's validated per request?)

func TestGameplay(t *testing.T) {
	createGame(t)
	addPlayer(t)
	attack(t)
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

func addPlayer(t *testing.T) {
	ship1 := bships.Ship{Start: bships.Coord{Row: 'B', Column: 1}, Dir: bships.TopToBottom, ShipType: mid}

	player := PlayerReq{PlayerName: "p1", Ships: []bships.Ship{ship1}}

	json, err := json.Marshal(player)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	req, _ := http.NewRequest("POST", "/bships/games/players", bytes.NewReader(json))
	resp := serve(req, PlayerHandler)

	checkResp(t, resp)
}

func attack(t *testing.T) {
	attack := AttackReq{PlayerName: "p1", Move: bships.Coord{Row: 'A', Column: 1}}

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
