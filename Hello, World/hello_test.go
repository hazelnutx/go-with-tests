package main

import "testing"

func TestHello(t *testing.T) {

	assetCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper is needed to be called, as when if the test fails, the error number
		// will be shown as it is in our function, rather than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Raul", "")
		want := "Hello, Raul"

		assetCorrectMessage(t, got, want)

	})

	t.Run("saying 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assetCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assetCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Raul", "French")
		want := "Bonjour, Raul"

		assetCorrectMessage(t, got, want)
	})
}
