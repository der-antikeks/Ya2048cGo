package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/der-antikeks/Ya2048cGo/game"
	"github.com/der-antikeks/Ya2048cGo/solvers"
)

type result struct {
	name  string
	score int
}

func main() {
	rand.Seed(time.Now().Unix())

	c := make(chan result)
	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		close(c)
	}()

	// TODO: timeout!
	for n, s := range solvers.List() {
		wg.Add(1)
		go func(name string, solver game.SolveFunc) {
			score, repts := 0, 1000
			for i := 0; i < repts; i++ {
				score += game.New().Run(solver)
			}
			c <- result{name, score / repts}
			wg.Done()
		}(n, s)
	}

	fmt.Println("Results:") // TODO: sorting?
	for r := range c {
		fmt.Printf("%s: %d\n", r.name, r.score)
	}
}
