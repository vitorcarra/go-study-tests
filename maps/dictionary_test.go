package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unkown")
		want := errNotFound

		assertError(t, want, got)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, errWordExists, err)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defintion := "definition"
		dictionary := Dictionary{word: defintion}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defintion := "definition"
		dictionary := Dictionary{word: defintion}
		newWord := "new"
		newDefinition := "new definition"

		err := dictionary.Update(newWord, newDefinition)

		assertError(t, errWordDoesNotExist, err)
	})
}

func TestDelete(t *testing.T) {

	t.Run("word exists", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)

		assertError(t, errNotFound, err)
	})

}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, definition, got)
}

func assertString(t testing.TB, want, got string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, want, got error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
