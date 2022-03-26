package types

import (
	"fmt"
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
	fmt.Printf("Philosopher %d take left forks\n", p.Id)
	p.left_fork.Take()
	fmt.Printf("Philosopher %d take right forks\n", p.Id)
	p.right_fork.Take()

	fmt.Printf("Philosopher %d is eating...\n", p.Id)
	time.Sleep(3 * time.Second)

	fmt.Printf("Philosopher %d put left forks\n", p.Id)
	p.left_fork.Put()
	fmt.Printf("Philosopher %d put right forks\n", p.Id)
	p.right_fork.Put()
}

func (p *Philosopher) Think() {
	fmt.Printf("Philosopher %d is thiking...\n", p.Id)
	time.Sleep(5 * time.Second)
}
