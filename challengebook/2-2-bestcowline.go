package challengebook

func bestCowLine(n int, s string) string {
	a := 0
	b := n - 1
	t := make([]byte, 0, n)
	for a <= b {
		var left bool
		for i := 0; a+i <= b; i++ {
			if s[a+i] < s[b-i] {
				left = true
				break
			} else if s[a+i] > s[b-i] {
				left = false
				break
			}
		}

		if left {
			t = append(t, s[a])
			a++
		} else {
			t = append(t, s[b])
			b--
		}
	}
	return string(t)
}
