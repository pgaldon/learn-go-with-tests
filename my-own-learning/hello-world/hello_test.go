package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello to somone in spanish", func(t *testing.T) {
		got := Hello("Jose", "SP")
		want := "Hola Jose"
		assertCorrectMessage(t, got, want)
	})
}
