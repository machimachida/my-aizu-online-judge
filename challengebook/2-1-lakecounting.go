package challengebook

func lakeCounting(n, m int, field [][]string) int {
	var count int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == "W" {
				checkAround(i, j, n, m, field)
				count++
			}
		}
	}
	return count
}

func checkAround(i, j, n, m int, field [][]string) {
	field[i][j] = "."

	for dn := -1; dn <= 1; dn++ {
		for dy := -1; dy <= 1; dy++ {
			if 0 <= i+dn && i+dn < n && 0 <= j+dy && j+dy < m && field[i+dn][j+dy] == "W" {
				checkAround(i+dn, j+dy, n, m, field)
			}
		}
	}
}
