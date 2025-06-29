package todo

import "sync"

var once = sync.Once{}

type Store struct {
	Todos map[string]*Todo
}

var store *Store

func GetStore() *Store {
	once.Do(func() {
		store = &Store{
			Todos: make(map[string]*Todo),
		}
	})
	return store
}
