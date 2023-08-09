package challengebook

func fenceRepair(n int, l []int) (ans int) {
	for n > 1 {
		mil1, mil2 := 0, 1
		if l[mil1] > l[mil2] {
			mil1, mil2 = mil2, mil1
		}

		for i := 2; i < n; i++ {
			if l[i] < l[mil1] {
				mil2 = mil1
				mil1 = i
			} else if l[i] < l[mil2] {
				mil2 = i
			}
		}

		t := l[mil1] + l[mil2]
		ans += t

		if mil1 == n-1 {
			mil1, mil2 = mil2, mil1
		}
		l[mil1] = t
		l[mil2] = l[n-1]
		n--
	}
	return ans
}
