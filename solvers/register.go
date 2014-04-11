package solvers

import (
	"sync"

	"github.com/der-antikeks/Ya2048cGo/game"
)

var (
	solvers map[string]game.SolveFunc
	lock    sync.Mutex
)

// Register a new Solver.
func Register(name string, f game.SolveFunc) {
	lock.Lock()
	defer lock.Unlock()

	if solvers == nil {
		solvers = make(map[string]game.SolveFunc)
	}
	solvers[name] = f
}

// Return a list of all registered Solvers.
func List() map[string]game.SolveFunc {
	lock.Lock()
	defer lock.Unlock()
	return solvers
}
