package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		_, err := dictionary.Search("unknown")
		want := ErrWordNotFound{word: "unknown"}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{
			word: definition,
		}
		err := dictionary.Add(word, "other definition")

		assertError(t, err, ErrWordExists{word})
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "a new definition"

		dict.Update(word, newDefinition)

		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("non existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}
		newDefinition := "a new definition"

		err := dict.Update(word, newDefinition)

		assertError(t, err, ErrWordNotFound{word})
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "this is just a test"}

	dict.Delete(word)

	_, err := dict.Search(word)
	assertError(t, err, ErrWordNotFound{word})
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
		t.Errorf("got error %q want %q", got, want)
	}

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("should find the added word %q but got error %q", word, err)
	}
	want := definition

	assertStrings(t, got, want)
}
