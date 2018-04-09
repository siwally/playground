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

func TestCreateGame(t *testing.T) {
	req, _ := http.NewRequest("POST", "/bships/games", nil)

	rec := serve(req, GameHandler)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}

	// TODO - Check gameId returned in JSON
	fmt.Printf("Create game returned body: %v\n", rec.Body)
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
	fmt.Printf("Add player returned body: %v\n", rec.Body)
}
func TestAttack(t *testing.T) {

	createTestGame()
	createTestPlayer(t)

	json := createAttackJSON(t)
	req, _ := http.NewRequest("POST", "/bships/games/attacks", bytes.NewReader(json))

	rec := serve(req, AttackHandler)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}

	// Check outcome in body
	fmt.Printf("Attack returned body: %v\n", rec.Body)
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

func createAttackJSON(t *testing.T) []byte {
	req := AttackReq{PlayerName: "p1", Move: bships.Coord{Row: 'A', Column: 1}}

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

func createTestPlayer(t *testing.T) {
	json := createPlayerJSON(t)
	req, _ := http.NewRequest("POST", "/bships/games/players", bytes.NewReader(json))

	serve(req, PlayerHandler)
}

func serve(req *http.Request, handlerFn func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFn)

	handler.ServeHTTP(rec, req)

	return rec
}
