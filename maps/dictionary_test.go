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
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("shound find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

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
