package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("passing a name returns greeting in english", func(t *testing.T) {
		got := Hello("Oliver", English)
		want := "Hello Oliver"
		assertCorrectMessage(t, got, want)
	})
	t.Run("passing an empty string returns default greeting", func(t *testing.T) {
		got := Hello("", English)
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("passing spanish arg returns Hola greeting", func(t *testing.T) {
		got := Hello("Senore", Spanish)
		want := "Hola Senore"
		assertCorrectMessage(t, got, want)
	})
	t.Run("passing french arg returns Bonjour greeting", func(t *testing.T) {
		got := Hello("Monsieur", French)
		want := "Bonjour Monsieur"
		assertCorrectMessage(t, got, want)
	})
}
