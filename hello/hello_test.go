package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Chris"

	if got := Hello("Chris"); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
