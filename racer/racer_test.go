package racer

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := makeDelayedServer(0)
	defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func BenchmarkRacer(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		delayStr := query.Get("delay")
		delay, err := strconv.Atoi(delayStr)
		if err != nil {
			delay = 0
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	urls := make([]string, 50)
	for i := 0; i < len(urls); i++ {
		urls[i] = server.URL + "?delay=" + strconv.Itoa(i+1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Racer(urls...)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
