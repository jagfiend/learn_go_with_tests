package select_keyword

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Racer method returns fastest result as expected", func(t *testing.T) {
		slowServer := makeServerWithDelay(20 * time.Millisecond)
		fastServer := makeServerWithDelay(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("got an unexpected error: %v", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Racer errors out if response too long", func(t *testing.T) {
		slowServer := makeServerWithDelay(25 * time.Millisecond)

		defer slowServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, slowServer.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expecting an error but didnt get one")
		}
	})
}

func makeServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
