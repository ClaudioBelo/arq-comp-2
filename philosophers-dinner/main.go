package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/ClaudioBelo/arq-comp-2/philosophers-dinner/types"
)

var forks = make([]*types.Fork, 5)
var philosophers = make([]*types.Philosopher, 5)

func main() {
	PhilosophersDinnerAsync()
}

func PhilosophersDinnerSync() {
	for i := 0; i < 5; i++ {
		forks[i] = types.NewFork(i)
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = types.NewPhilosopher(i, forks[i], forks[(i+1)%5])
	}
	for {
		if err := philosophers[rand.Intn(4)].Next(); err != nil {
			fmt.Printf("%s", err)
		}
	}
}

func PhilosophersDinnerAsync() {
	for i := 0; i < 5; i++ {
		forks[i] = types.NewFork(i)
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
