package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	score := s.scores[name]
	s.scores[name] = score + 1
}

func TestPOSTPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		postRequest := NewPostScoreRequest("Pepper")
		postResponse := httptest.NewRecorder()

		server.ServeHTTP(postResponse, postRequest)

		AssertStatus(t, postResponse.Code, http.StatusAccepted)

		getRequest := NewGetScoreRequest("Pepper")
		getResponse := httptest.NewRecorder()
		server.ServeHTTP(getResponse, getRequest)

		AssertScore(t, getResponse.Body.String(), "1")
	})
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertScore(t, response.Body.String(), "20")
		AssertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request := NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertScore(t, response.Body.String(), "10")
		AssertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		AssertStatus(t, got, want)
	})
}

func AssertStatus(t testing.TB, got int, want int) {
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func NewPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
	return req
}

func AssertScore(t testing.TB, got string, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
