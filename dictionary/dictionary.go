package dictionary

import (
	"fmt"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]
	if !found {
		return "", fmt.Errorf("could not find the word %q", word)
	}

	return definition, nil
}
