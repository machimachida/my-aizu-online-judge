package challengebook

import "testing"

func TestSubsetSum(t *testing.T) {
	t.Run("subsetSum(4, 13, []int{1, 2, 4, 7})", func(t *testing.T) {
		if !subsetSum(4, 13, []int{1, 2, 4, 7}) {
			t.Error("got false, want true")
		}
	})
	t.Run("subsetSum(4, 15, []int{1, 2, 4, 7})", func(t *testing.T) {
		if subsetSum(4, 15, []int{1, 2, 4, 7}) {
			t.Error("got true, want false")
		}
	})

}
