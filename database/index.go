package database

import (
	"errors"
	"sync"
)

type KeyStoreMap map[string]string

type Store struct {
	mu     sync.Mutex
	keyMap map[string]string
}

func CreateNewStore() *Store {
	return &Store{
		keyMap: make(map[string]string),
	}
}

func (store *Store) Get(key string) (string, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	val, exists := store.keyMap[key]
	if !exists {
		return "", errors.New("value not present in the store")
	}
	return val, nil
}

func (store *Store) Set(key string, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.keyMap[key] = value
}

func (store *Store) Remove(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	keyMap := store.keyMap
	delete(keyMap, key)

	store.keyMap = keyMap
}
