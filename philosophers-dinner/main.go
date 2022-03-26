package main

import (
	"sync"

	"github.com/ClaudioBelo/arq-comp-2/philosophers-dinner/types"
)

var forks = make([]*types.Fork, 5)
var philosophers = make([]*types.Philosopher, 5)

func main() {
	PhilosophersDinnerSync()
}

func PhilosophersDinnerSync() {
	for i := 0; i < 5; i++ {
		forks[i] = types.NewFork()
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = types.NewPhilosopher(i, forks[i], forks[(i+1)%5])
	}

	for {
		for i := 0; i < 5; i++ {
			philosophers[i].Think()
			philosophers[i].Eat()
		}
	}
}

func PhilosophersDinnerAsync() {
	for i := 0; i < 5; i++ {
		forks[i] = types.NewFork()
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = types.NewPhilosopher(i, forks[i], forks[(i+1)%5])
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].Exists()
	}
	wg.Wait()
}
