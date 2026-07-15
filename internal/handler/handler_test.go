package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seuusername/guess-game/internal/store"
)

func TestCreateGame(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	req := httptest.NewRequest(http.MethodPost, "/game", nil)
	w := httptest.NewRecorder()

	h.CreateGame(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}

	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	id, ok := response["id"]
	if !ok || id == "" {
		t.Error("Expected id in response")
	}
}

func TestCreateGame_WrongMethod(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	req := httptest.NewRequest(http.MethodGet, "/game", nil)
	w := httptest.NewRecorder()

	h.CreateGame(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", w.Code)
	}
}

func TestGuessNumber_Success(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	// Create a game first
	req := httptest.NewRequest(http.MethodPost, "/game", nil)
	w := httptest.NewRecorder()
	h.CreateGame(w, req)

	var createResponse map[string]string
	json.NewDecoder(w.Body).Decode(&createResponse)
	gameID := createResponse["id"]

	// Make a guess
	guessBody := map[string]int{"number": 50}
	bodyBytes, _ := json.Marshal(guessBody)
	req = httptest.NewRequest(http.MethodPost, "/game/"+gameID+"/guess", bytes.NewReader(bodyBytes))
	req.SetPathValue("id", gameID)
	w = httptest.NewRecorder()

	h.GuessNumber(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)

	if _, ok := response["correct"]; !ok {
		t.Error("Expected 'correct' in response")
	}
	if _, ok := response["attempts"]; !ok {
		t.Error("Expected 'attempts' in response")
	}
}

func TestGuessNumber_GameNotFound(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	guessBody := map[string]int{"number": 50}
	bodyBytes, _ := json.Marshal(guessBody)
	req := httptest.NewRequest(http.MethodPost, "/game/non-existent/guess", bytes.NewReader(bodyBytes))
	req.SetPathValue("id", "non-existent")
	w := httptest.NewRecorder()

	h.GuessNumber(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

func TestGuessNumber_InvalidBody(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	// Create a game first
	req := httptest.NewRequest(http.MethodPost, "/game", nil)
	w := httptest.NewRecorder()
	h.CreateGame(w, req)

	var createResponse map[string]string
	json.NewDecoder(w.Body).Decode(&createResponse)
	gameID := createResponse["id"]

	// Send invalid JSON
	req = httptest.NewRequest(http.MethodPost, "/game/"+gameID+"/guess", bytes.NewReader([]byte("invalid json")))
	req.SetPathValue("id", gameID)
	w = httptest.NewRecorder()

	h.GuessNumber(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestGetGameStatus_Success(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	// Create a game
	req := httptest.NewRequest(http.MethodPost, "/game", nil)
	w := httptest.NewRecorder()
	h.CreateGame(w, req)

	var createResponse map[string]string
	json.NewDecoder(w.Body).Decode(&createResponse)
	gameID := createResponse["id"]

	// Get status
	req = httptest.NewRequest(http.MethodGet, "/game/"+gameID, nil)
	req.SetPathValue("id", gameID)
	w = httptest.NewRecorder()

	h.GetGameStatus(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)

	if attempts, ok := response["attempts"].(float64); !ok || attempts != 0 {
		t.Errorf("Expected 0 attempts, got %v", response["attempts"])
	}
}

func TestGetGameStatus_NotFound(t *testing.T) {
	s := store.NewStore()
	h := NewHandler(s)

	req := httptest.NewRequest(http.MethodGet, "/game/non-existent", nil)
	req.SetPathValue("id", "non-existent")
	w := httptest.NewRecorder()

	h.GetGameStatus(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}
