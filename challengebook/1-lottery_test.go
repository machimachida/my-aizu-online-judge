package challengebook

import "testing"

func TestLottery(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		n := 3
		m := 10
		k := []int{1, 3, 5}
		if !lottery(n, m, k) {
			t.Errorf("lottery(%d, %d, %v) should be true, but false", n, m, k)
		}
	})

	t.Run("case2", func(t *testing.T) {
		n := 3
		m := 9
		k := []int{1, 3, 5}
		if lottery(n, m, k) {
			t.Errorf("lottery(%d, %d, %v) should be false, but true", n, m, k)
		}
	})
}
