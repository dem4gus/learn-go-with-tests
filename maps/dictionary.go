package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]
	if !ok {
		return "", errors.New("could not find term")
	}
	return definition, nil
}
