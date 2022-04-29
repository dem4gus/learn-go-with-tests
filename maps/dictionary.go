package dictionary

const (
	ErrNotFound         = DictionaryErr("could not find term")
	ErrWordExists       = DictionaryErr("word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("could not update word, does not exist in dictionary")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(term, definition string) error {
	_, err := d.Search(term)
	switch err {
	case nil:
		d[term] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}
