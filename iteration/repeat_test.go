package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertRepeat := func(t testing.TB, expected, result string) {
		t.Helper()
		if result != expected {
			t.Errorf("expected %q but got %q", expected, result)
		}
	}
	t.Run("test single character", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeat(t, expected, repeated)
	})
	t.Run("test single repetition", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"
		assertRepeat(t, expected, repeated)
	})
	t.Run("test single repetition of multiple characters", func(t *testing.T) {
		repeated := Repeat("abc", 1)
		expected := "abc"
		assertRepeat(t, expected, repeated)
	})
	t.Run("test multiple repetitions of multiple characters", func(t *testing.T) {
		repeated := Repeat("AbC", 3)
		expected := "AbCAbCAbC"
		assertRepeat(t, expected, repeated)
	})
	t.Run("test repetition with rune", func(t *testing.T) {
		repeated := Repeat("💩", 3)
		expected := "💩💩💩"
		assertRepeat(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
