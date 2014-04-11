package game

import (
	"fmt"
	"math/rand"
)

type SolveFunc func(grid []int) Direction

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Yet another 2048 clone
type Game struct {
	grid  []int
	score int
}

// Creates a new game.
func New() *Game {
	g := &Game{make([]int, 16), 0}
	g.add()
	return g
}

func (g *Game) add() bool {
	empty := []int{}
	for p, v := range g.grid {
		if v == 0 {
			empty = append(empty, p)
		}
	}

	if l := len(empty); l > 0 {
		g.grid[empty[rand.Intn(l)]] = 2
		return true
	}

	return false
}

// Returns a tab delimited representation of the game grid.
func (g *Game) String() string {
	s := ""
	for p, v := range g.grid {
		if p > 0 && p%4 == 0 {
			s += "\n"
		}
		s += fmt.Sprintf(" % 4d", v)
	}
	return s + "\n"
}

// Runs a given solving function until there are no more options available
// and returns the final score.
func (g *Game) Run(f SolveFunc) (score int) {
	attempts := 0
	for g.check() {
		s, ok := g.move(f(g.grid))
		if ok {
			attempts = 0
			if !g.add() {
				return
			}
		} else {
			if attempts++; attempts > 10 {
				return
			}
		}
		score += s
	}
	return
}

func (g *Game) move(d Direction) (score int, moved bool) {
	var z, f, l int

	switch d {
	case Up:
		z, f, l = 4, 4, 1
	case Down:
		z, f, l = 11, -4, -1
	case Left:
		z, f, l = 1, 1, 4
	case Right:
		z, f, l = 14, -1, -4
	}

	for s := 0; s < 12; s++ {
		c := z + (s % 3 * f) + (s / 3 * l)
		if g.grid[c] == 0 {
			continue
		}

		min := z - f + (s / 3 * l)
		n := c
		for {
			n -= f

			if g.grid[n] != 0 {
				if g.grid[n] != g.grid[c] {
					n += f
				}
				break
			}

			if n == min {
				break
			}
		}

		if c == n {
			continue
		}

		if g.grid[n] == g.grid[c] {
			g.grid[n] += g.grid[c]
			g.grid[c] = 0
			moved = true
			score += g.grid[n]
		} else if g.grid[n] == 0 {
			g.grid[n] = g.grid[c]
			g.grid[c] = 0
			moved = true
		}
	}

	return
}

func (g *Game) check() bool {
	for s := 0; s < 16; s++ {
		c := g.grid[s]

		if c == 0 {
			return true
		}

		if r := (s + 1) % 4; r > 0 {
			if c == g.grid[r+(s/4)*4] {
				return true
			}
		}

		if s < 12 {
			if c == g.grid[s+4] {
				return true
			}
		}
	}

	return false
}
