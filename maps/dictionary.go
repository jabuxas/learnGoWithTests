package main

type Dictionary map[string]string

const (
	ErrNotFound          = DictionaryErr("the word is not in the dictionary")
	ErrWordExists        = DictionaryErr("the word is already in the dictionary")
	ErrWordDoesNotExists = DictionaryErr("the word doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value

	case nil:
		return ErrWordExists

	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists

	case nil:
		d[word] = definition
		return nil

	}
	return nil
}
