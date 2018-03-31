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
	json := createGameJSON(t, 10, 10)
	req, _ := http.NewRequest("POST", "/bships/games", bytes.NewReader(json))

	rec := serve(req, GameHandler)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}
}

func TestAddPlayer(t *testing.T) {
	json := createPlayerJSON(t)
	req, _ := http.NewRequest("POST", "/bships/games/players", bytes.NewReader(json))

	rec := serve(req, PlayerHandler)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v instead of %v", status, http.StatusOK)
	}
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

func createGameJSON(t *testing.T, gridWidth, gridHeight int) []byte {
	cfg := bships.GameConfig{GridWidth: 10, GridHeight: 10}
	json, err := json.Marshal(cfg)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	return json
}

func createPlayerJSON(t *testing.T) []byte {

	req := PlayerReq{PlayerName: "p1"} // , Ships: []bships.Ship{}

	json, err := json.Marshal(req)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	return json
}

func serve(req *http.Request, handlerFn func(http.ResponseWriter, *http.Request)) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerFn)

	handler.ServeHTTP(rec, req)

	return rec
}
