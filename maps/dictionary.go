package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find term")
var ErrWordExists = errors.New("word already exists in dictionary")

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)
	switch err {
	case ErrNotFound:
		d[term] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
