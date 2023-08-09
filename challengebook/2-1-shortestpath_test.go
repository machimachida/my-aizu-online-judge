package challengebook

import (
	"strings"
	"testing"
)

func TestShortestPath(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		n := 10
		m := 10
		field := make([][]string, n)
		field[0] = strings.Split("#S######.#", "")
		field[1] = strings.Split("......#..#", "")
		field[2] = strings.Split(".#.##.##.#", "")
		field[3] = strings.Split(".#........", "")
		field[4] = strings.Split("##.##.####", "")
		field[5] = strings.Split("....#....#", "")
		field[6] = strings.Split(".#######.#", "")
		field[7] = strings.Split("....#.....", "")
		field[8] = strings.Split(".####.###.", "")
		field[9] = strings.Split("....#...G#", "")
		out := shortestPath(n, m, field)
		if out != 22 {
			t.Errorf("got %v, want %v", out, 22)
		}
	})
}
