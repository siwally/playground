package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/siwally/golang/bships"
)

func TestCreateGame(t *testing.T) {
	req, _ := http.NewRequest("POST", "/bships/games", nil)

	rec := serve(req, GameHandler)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}

	// TODO - Check gameId returned in JSON
}

func TestAddPlayer(t *testing.T) {

	createTestGame() // TODO - use game id returned

	json := createPlayerJSON(t)
	req, _ := http.NewRequest("POST", "/bships/games/players", bytes.NewReader(json))

	rec := serve(req, PlayerHandler)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}

	// TODO - Check playerId returned in JSON
}
func TestMove(t *testing.T) {
	// TODO - utilities to create game, create player, using ids, then make a move and validate
}

func DontTestServer(t *testing.T) {

	go StartServer(":8080")

	resp, err := http.Post("http://localhost:8080/bships/game", "application/json", nil)

	if err != nil {
		t.Errorf("Failed to post; err is %v", err)
	}

	resp.Write(os.Stdout)
}

// Small utitlity methods to simplify tests.

func createPlayerJSON(t *testing.T) []byte {

	ship1 := bships.Ship{Start: bships.Coord{Row: 'B', Column: 1}, Dir: bships.TopToBottom, ShipType: mid}

	req := PlayerReq{PlayerName: "p1", Ships: []bships.Ship{ship1}}

	json, err := json.Marshal(req)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	return json
}

func createTestGame() {
	req, _ := http.NewRequest("POST", "/bships/games", nil)

	serve(req, GameHandler)
}

func serve(req *http.Request, handlerFn func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFn)

	handler.ServeHTTP(rec, req)

	return rec
}
