package types

import "errors"

type Fork struct {
	Id     int
	in_use bool
}

func NewFork(id int) *Fork {
	return &Fork{
		Id:     id,
		in_use: false,
	}
}

func (f *Fork) Take() error {
	if !f.in_use {
		f.in_use = true
		return nil
	} else {
		return errors.New("Fork is already in use")
	}
}

func (f *Fork) Put() {
	f.in_use = false
}
