package challengebook

import (
	"testing"
)

func TestSarumansArmy(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := 6
		r := 10
		x := []int{1, 7, 15, 20, 30, 50}
		want := 3
		got := salmansArmy(n, r, x)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
