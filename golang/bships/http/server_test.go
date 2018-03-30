package http

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateGame(t *testing.T) {

	// TODO Convert GameConfig to JSON and post data
	// cfg := bships.GameConfig{GridHeight: 10, GridWidth: 10}

	// bytes, _ := json.Marshal(cfg)

	req, err := http.NewRequest("POST", "http://localhost:8080/bships/game", nil)

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

func TestServer(t *testing.T) {

	go StartServer(":8080")

	resp, err := http.Post("http://localhost:8080/bships/game", "application/json", nil)

	if err != nil {
		t.Errorf("Failed to post; err is %v", err)
	}

	resp.Write(os.Stdout)
}
