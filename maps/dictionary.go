package maps

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("could not find the word you were looking for")
	errWordExists = errors.New("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	default:
		return err
	}

	return nil
}
