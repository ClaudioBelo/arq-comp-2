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
}

func NewPhilosopher(id int, left_fork, right_fork *Fork) *Philosopher {
	return &Philosopher{
		Id:         id,
		left_fork:  left_fork,
		right_fork: right_fork,
	}
}

func (p *Philosopher) Eat() {
	fmt.Printf("Philosopher %d take left and right fork\n", p.Id)
	p.left_fork.Take()
	p.right_fork.Take()

	fmt.Printf("Philosopher %d is eating...\n", p.Id)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

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
		p.Eat()
	}
}
