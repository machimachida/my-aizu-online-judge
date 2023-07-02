package challengebook

func subsetSum(n, k int, a []int) bool {
	return depthFirstSearch(0, 0, k, a)
}

func depthFirstSearch(i, sum, k int, a []int) bool {
	if i == len(a) {
		return sum == k
	}
	if depthFirstSearch(i+1, sum, k, a) {
		return true
	}
	if depthFirstSearch(i+1, sum+a[i], k, a) {
		return true
	}
	return false
}
