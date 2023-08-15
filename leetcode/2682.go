package lc

func circularGameLosers(n int, k int) []int {
	flags := make([]bool, n)
	tempk := k
	nextN := 0
	for {
		if flags[nextN] {
			break
		}
		flags[nextN] = true
		nextN = (nextN + tempk) % n
		tempk = (tempk + k) % n
	}
	ans := make([]int, 0)
	for inde, v := range flags {
		if !v {
			ans = append(ans, inde+1)
		}
	}
	return ans
}
