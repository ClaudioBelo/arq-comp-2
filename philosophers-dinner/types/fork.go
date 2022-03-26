package types

import (
	"sync"
)

type Fork struct {
	mu sync.Mutex
}

func NewFork() *Fork {
	return &Fork{
		mu: sync.Mutex{},
	}
}

func (f *Fork) Take() {
	f.mu.Lock()
}

func (f *Fork) Put() {
	f.mu.Unlock()
}
