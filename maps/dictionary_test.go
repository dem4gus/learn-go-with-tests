package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, "this is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("foo")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected error")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		term := "test"
		definition := "this is just a test"
		dictionary.Add("test", "this is just a test")
		assertDefinition(t, dictionary, term, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}
		err := dictionary.Add(term, "this is a new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, term, definition)
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, term, definition string) {
	t.Helper()
	got, err := dictionary.Search(term)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
