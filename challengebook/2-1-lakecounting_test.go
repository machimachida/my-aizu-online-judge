package challengebook

import (
	"testing"
)

func TestLakeCounting(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		n := 10
		m := 12
		field := make([][]string, n)
		for i := 0; i < n; i++ {
			field[i] = make([]string, m)
			for j := 0; j < m; j++ {
				field[i][j] = "."
			}
		}
		field[0][0] = "W"
		field[0][9] = "W"
		field[0][10] = "W"

		field[1][1] = "W"
		field[1][2] = "W"
		field[1][3] = "W"
		field[1][9] = "W"
		field[1][10] = "W"
		field[1][11] = "W"

		field[2][4] = "W"
		field[2][5] = "W"
		field[2][9] = "W"
		field[2][10] = "W"

		field[3][9] = "W"
		field[3][10] = "W"

		field[4][9] = "W"

		field[5][2] = "W"
		field[5][9] = "W"

		field[6][1] = "W"
		field[6][3] = "W"
		field[6][9] = "W"
		field[6][10] = "W"

		field[7][0] = "W"
		field[7][2] = "W"
		field[7][4] = "W"
		field[7][10] = "W"

		field[8][1] = "W"
		field[8][3] = "W"
		field[8][10] = "W"

		field[9][2] = "W"
		field[9][10] = "W"

		count := lakeCounting(n, m, field)
		if count != 3 {
			t.Errorf("got %v, want %v", count, 3)
		}
	})
}
