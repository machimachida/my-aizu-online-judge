package challengebook

func selectCoin(coins []int, a int) int {
	coinVals := []int{1, 5, 10, 50, 100, 500}
	ans := 0
	for i := len(coins) - 1; i >= 0; i-- {
		var t int
		if n := a / coinVals[i]; n < coins[i] {
			t = n
		} else {
			t = coins[i]
		}
		a -= t * coinVals[i]
		ans += t
	}
	return ans
}
