package challengebook

import (
	"sort"
)

func lottery(n, m int, k []int) bool {
	kk := make([]int, n*n)
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			kk[a*n+b] = k[a] + k[b]
		}
	}
	sort.Ints(kk)

	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			diff := m - k[a] - k[b]
			index := sort.SearchInts(kk, diff)
			if index < len(kk) && kk[index] == diff {
				return true
			}
		}
	}
	return false
}
