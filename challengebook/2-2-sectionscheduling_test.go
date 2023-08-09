package challengebook

import (
	"testing"
)

func TestSectionScheduling(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := 5
		start := []int{1, 2, 4, 6, 8}
		end := []int{3, 5, 7, 9, 10}
		want := 3
		got := sectionScheduling(n, start, end)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
