package solvers

import (
	"github.com/der-antikeks/Ya2048cGo/game"
)

func init() {
	dirs := []game.Direction{game.Left, game.Down, game.Up}
	step := 0

	Register("Repeating Solver", func(grid []int) game.Direction {
		step++
		return dirs[step%3]
	})
}
