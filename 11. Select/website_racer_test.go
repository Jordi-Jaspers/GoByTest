package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("When a slow and fast url are given, the fast url is returned", func(t *testing.T) {
		slowServer := CreateServerWithStatus(20*time.Millisecond, http.StatusOK)
		fastServer := CreateServerWithStatus(0, http.StatusOK)

		want := fastServer.URL
		got, _ := Racer(slowServer.URL, fastServer.URL)
		AssertUrls(t, got, want)

		slowServer.Close()
		fastServer.Close()
	})
	t.Run("When none of the urls return within 10 seconds then it should return an error", func(t *testing.T) {
		slowServer := CreateServerWithStatus(12*time.Second, http.StatusOK)
		fastServer := CreateServerWithStatus(12*time.Second, http.StatusOK)

		_, err := ConfigurableRacer(slowServer.URL, fastServer.URL, 0)
		AssertError(t, err)

		slowServer.Close()
		fastServer.Close()
	})
}

func CreateServerWithStatus(delay time.Duration, status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(status)
	}))
}

func AssertError(t testing.TB, err error) {
	if err == nil {
		t.Fatal("Expected to receive an error")
	}
	if err.Error() != noResponseException {
		t.Errorf("got %q, want %q \n", err.Error(), noResponseException)
	}
}

func AssertUrls(t testing.TB, got string, want string) {
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
