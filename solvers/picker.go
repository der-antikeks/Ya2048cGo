package solvers

import (
	"github.com/der-antikeks/Ya2048cGo/game"
	"math/rand"
)

func init() {
	dirs := []game.Direction{game.Up, game.Down, game.Left, game.Right}

	Register("Cherry Picker", func(grid []int) game.Direction {
		scores := make([]int, len(dirs))
		found := 0

		for s := 0; s < 16; s++ {
			c := grid[s]

			if c == 0 {
				continue
			}

			// right
			if r := (s + 1) % 4; r > 0 {
				if c == grid[r+(s/4)*4] {
					scores[3] += c
					found++
				}
			}

			// left
			if l := s % 4; l > 0 {
				if c == grid[l+((s/4)*4)-1] {
					scores[2] += c
					found++
				}
			}

			// bottom
			if s < 12 {
				if c == grid[s+4] {
					scores[1] += c
					found++
				}
			}

			// top
			if s > 3 {
				if c == grid[s-4] {
					scores[0] += c
					found++
				}
			}
		}

		pick, max := 0, 0
		for d, s := range scores {
			if s > max {
				max = s
				pick = d
			}
		}

		if found == 0 {
			pick = rand.Intn(4)
		}

		return dirs[pick]
	})
}
