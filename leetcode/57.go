package lc

func insert(intervals [][]int, newInterval []int) [][]int {
	lenI := len(intervals)
	ans := make([][]int, 0)
	i := 0
	start, end := newInterval[0], newInterval[1]
	for ; i < lenI; i++ {
		if intervals[i][1] < start {
			ans = append(ans, intervals[i])
		} else {
			if intervals[i][0] < start {
				start = intervals[i][0]
			}
			break
		}
	}
	for ; i < lenI; i++ {
		if intervals[i][0] > end {
			break
		}
	}
	if i-1 >= 0 {
		if intervals[i-1][1] > end {
			end = intervals[i-1][1]
		}
	}
	ans = append(ans, []int{start, end})
	for ; i < lenI; i++ {
		ans = append(ans, intervals[i])
	}
	return ans
}
