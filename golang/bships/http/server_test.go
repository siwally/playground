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

	cfg := bships.GameConfig{GridWidth: 10, GridHeight: 10}
	json, err := json.Marshal(cfg)

	if err != nil {
		t.Fatalf("Unable to encode request data for test: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/bships/game", bytes.NewReader(json))

	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(GameHandler)

	handler.ServeHTTP(rec, req)

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
