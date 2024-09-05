package dictionary

import (
	"fmt"
)

type ErrWordNotFound struct {
	word string
}

func (err ErrWordNotFound) Error() string {
	return fmt.Sprintf("could not find the word %q", err.word)
}

type ErrWordExists struct {
	word string
}

func (err ErrWordExists) Error() string {
	return fmt.Sprintf("word %q already has a definition", err.word)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", ErrWordNotFound{word}
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, wordAlreadyExists := d[word]
	if wordAlreadyExists {
		return ErrWordExists{word}
	}

	d[word] = definition
	return nil
}
