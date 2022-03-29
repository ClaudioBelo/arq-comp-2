package types

import (
	"fmt"
	"sync"
)

type Fork struct {
	Id     int
	mu     sync.Mutex
	locked bool
}

func NewFork(id int) *Fork {
	return &Fork{
		Id:     id,
		mu:     sync.Mutex{},
		locked: false,
	}
}

func (f *Fork) Take() error {
	if f.locked {
		return fmt.Errorf("Fork %d in use\n", f.Id)
	}
	f.locked = true
	f.mu.Lock()

	return nil
}

func (f *Fork) Put() {
	f.locked = false
	f.mu.Unlock()
}
