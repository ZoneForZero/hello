package lc

import "fmt"

func summaryRanges(nums []int) []string {
	lenN := len(nums)
	if lenN == 0 {
		return []string{}
	}
	ans := []string{}
	sI := 0
	for i := 1; i < lenN; i++ {
		if nums[i]-nums[i-1] != 1 {
			if sI != i-1 {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[sI], nums[i-1]))
			} else {
				ans = append(ans, fmt.Sprintf("%d", nums[sI]))
			}
			sI = i
		}
	}
	if sI != lenN-1 {
		ans = append(ans, fmt.Sprintf("%d->%d", nums[sI], nums[lenN-1]))
	} else {
		ans = append(ans, fmt.Sprintf("%d", nums[sI]))
	}
	return ans
}
