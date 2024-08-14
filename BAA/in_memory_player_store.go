package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
	if _, ok := i.store[name]; !ok {
		return 0, ok
	}
	return i.store[name], true
}
