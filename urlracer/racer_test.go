package urlracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare two servers", func(t *testing.T) {
		slowURL := makeDelayedServer(20 * time.Millisecond)
		fastURL := makeDelayedServer(0 * time.Millisecond)

		defer slowURL.Close()
		defer fastURL.Close()

		want := fastURL.URL
		got, _ := Racer(slowURL.URL, fastURL.URL)

		if want != got {
			t.Errorf("want url %q but got url %q", want, got)
		}
		slowURL.Close()
		fastURL.Close()
	})

	t.Run("return timeout after 10s", func(t *testing.T) {
		timeout := 20 * time.Millisecond
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		if err == nil {
			t.Errorf("it was expecting a timeout error")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
