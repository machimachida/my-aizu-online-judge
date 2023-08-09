package challengebook

func salmansArmy(n, r int, x []int) (ans int) {
	var s int
	for i := 0; i < n; {
		s = x[i]

		i++
		for i < n && x[i] <= s+r {
			i++
		}

		p := x[i-1]

		for i < n && x[i] <= p+r {
			i++
		}
		ans++
	}
	return
}
