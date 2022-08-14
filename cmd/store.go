package main

import "errors"

func Put(key string, value string) error {
	store[key] = value

	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}

var ErrorNoSuchKey = errors.New("No such key")

var store = make(map[string]string)
