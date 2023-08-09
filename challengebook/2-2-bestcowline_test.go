package challengebook

import (
	"testing"
)

func TestBestCowLine(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := 6
		s := "ACDBCB"
		want := "ABCBCD"
		got := bestCowLine(n, s)
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
