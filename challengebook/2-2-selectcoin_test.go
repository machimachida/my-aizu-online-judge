package challengebook

import (
	"testing"
)

func TestSelectCoin(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		coins := []int{3, 2, 1, 3, 0, 2}
		a := 620
		want := 6
		got := selectCoin(coins, a)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
