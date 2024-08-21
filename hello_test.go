package main

import "testing"

func TestHello_World(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello_Guilherme(t *testing.T) {
	got := Hello("Guilherme")
	want := "Hello, Guilherme!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
