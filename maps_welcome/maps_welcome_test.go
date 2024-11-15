package maps_welcome

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	assertStringsMatch := func(t testing.TB, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertDefinition := func(t testing.TB, dictionary Dictionary, word string, definition string) {
		t.Helper()

		got, err := dictionary.Search(word)

		if err != nil {
			t.Fatal(err)
		}

		assertStringsMatch(t, got, definition)
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("can add a word to dictionary as expected", func(t *testing.T) {
		word := "cat"
		defintion := "fluffy wonder"

		dictionary := Dictionary{}

		err := dictionary.Add(word, defintion)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, defintion)
	})

	t.Run("wont override existing definition", func(t *testing.T) {
		word := "cat"
		definition := "fluffy wonder"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "nothing but hairballs")

		assertError(t, err, ErrWordAlreadyDefined)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("finds given word from dictionary as expected", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStringsMatch(t, got, want)
	})

	t.Run("returns error if given word not found", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrWordNotFound)
	})

	t.Run("can update an existing definition", func(t *testing.T) {
		word := "cat"
		definition := "fluffy wonder"
		dictionary := Dictionary{word: definition}
		newDefinition := "a rare beauty"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("will error if no word found to update", func(t *testing.T) {
		word := "cat"
		definition := "fluffy wonder"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesntExist)
	})

	t.Run("will delete as expected", func(t *testing.T) {
		word := "cat"
		dictionary := Dictionary{word: "fluffy wonder"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrWordNotFound)
	})
}
