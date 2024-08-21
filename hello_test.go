package main

import "testing"

func TestHello_World(t *testing.T) {
	t.Run("saying hello to world when a empty string is supplied", func(t *testing.T) {
		got := Hello()
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Guilherme", func(t *testing.T) {
		got := Hello("Guilherme")
		want := "Hello, Guilhermee!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Guilherme Teixeira", func(t *testing.T) {
		got := Hello("Guilherme", "Teixeira")
		want := "Hello, Guilherme Teixeira!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
