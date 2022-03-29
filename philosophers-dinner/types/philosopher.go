package types

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	Id         int
	left_fork  *Fork
	right_fork *Fork
	state      int
}

func NewPhilosopher(id int, left_fork, right_fork *Fork) *Philosopher {
	return &Philosopher{
		Id:         id,
		left_fork:  left_fork,
		right_fork: right_fork,
		state:      0,
	}
}

func (p *Philosopher) Next() error {
	switch p.state {
	case 0:
		fmt.Printf("Philosopher %d take left fork(%d)\n", p.Id, p.left_fork.Id)
		if err := p.left_fork.Take(); err != nil {
			return err
		}
		p.state = 1
	case 1:
		fmt.Printf("Philosopher %d take right fork(%d)\n", p.Id, p.right_fork.Id)
		if err := p.right_fork.Take(); err != nil {
			return err
		}
		p.state = 2
	case 2:
		p.Eat()
		p.LetForks()
		p.state = 0
	}

	return nil
}

func (p *Philosopher) TakeForks() error {
	fmt.Printf("Philosopher %d take left and right fork\n", p.Id)
	var err error
	if err = p.left_fork.Take(); err != nil {
		return err
	}
	if err = p.right_fork.Take(); err != nil {
		return err
	}
	return nil
}

func (p *Philosopher) Eat() {
	fmt.Printf("Philosopher %d is eating...\n", p.Id)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}

func (p *Philosopher) LetForks() {
	fmt.Printf("Philosopher %d put left and right fork\n", p.Id)
	p.left_fork.Put()
	p.right_fork.Put()
}

func (p *Philosopher) Think() {
	fmt.Printf("Philosopher %d is thiking...\n", p.Id)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}

func (p *Philosopher) Exists() {
	for {
		p.Think()
		if err := p.TakeForks(); err != nil {
			fmt.Printf("%s", err)
			continue
		}
		p.Eat()
		p.LetForks()
	}
}
