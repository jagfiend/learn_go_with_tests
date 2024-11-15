package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test dependency injection", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Wondini")

		got := buffer.String()
		want := "Hello Wondini"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
