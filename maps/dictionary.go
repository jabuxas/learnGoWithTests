package main

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("the word is already in the dictionary")
)

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
