package main

import (
	"errors"
	"sync"
)

func Put(key string, value string) error {
	store.Lock()
	store.m[key] = value
	store.Unlock()

	return nil
}

func Get(key string) (string, error) {
	store.RLock()
	value, ok := store.m[key]
	store.RUnlock()
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	store.Lock()
	delete(store.m, key)
	store.Unlock()
	return nil
}

var ErrorNoSuchKey = errors.New("No such key")

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

//make(map[string]string)
