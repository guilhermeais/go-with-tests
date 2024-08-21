package main

import "testing"

func TestHello_World(t *testing.T) {
	t.Run("saying hello to world when a empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Guilherme", func(t *testing.T) {
		got := Hello("Guilherme", "")
		want := "Hello, Guilherme!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Guilherme in Portuguese", func(t *testing.T) {
		got := Hello("Guilherme", Portuguese)
		want := "Olá, Guilherme!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Guilherme in Spanish", func(t *testing.T) {
		got := Hello("Guilherme", "Spanish")
		want := "Hola, Guilherme!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
