package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("adding two integers returns expected integer", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 7
}
