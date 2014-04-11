package solvers

import (
	"math/rand"

	"github.com/der-antikeks/Ya2048cGo/game"
)

func init() {
	Register("Random Solver", NewRandom().Solve)
}

type Random struct {
	steps int
	dirs  []game.Direction
}

func NewRandom() *Random {
	return &Random{
		dirs: []game.Direction{game.Up, game.Down, game.Left, game.Right},
	}
}

func (s *Random) Solve(grid []int) game.Direction {
	s.steps++
	return s.dirs[rand.Intn(4)]
}
