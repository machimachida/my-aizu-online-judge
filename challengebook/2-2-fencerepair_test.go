package challengebook

import (
	"testing"
)

func TestFenceRepair(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := 3
		l := []int{8, 5, 8}
		want := 34
		got := fenceRepair(n, l)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		n := 4
		l := []int{4, 4, 4, 4}
		want := 32
		got := fenceRepair(n, l)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
