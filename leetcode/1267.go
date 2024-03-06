package lc

func countServers(grid [][]int) int {
	return 0
}
func furthestDistanceFromOrigin(moves string) int {
	gn := 0
	w := 0
	//maxv := 0
	for _, v := range moves {
		if v == '_' {
			gn++
		} else if v == 'L' {
			w--
		} else {
			w++
		}
		//if w+gn > maxv {
		//	maxv = w + gn
		//}
		//if (w-gn)*(-1) > maxv {
		//	maxv = (w - gn) * -1
		//}
	}
	if w > 0 {
		return w + gn
	}
	return (w * -1) + gn
	//return maxv
}

func minimumPossibleSum(n int, target int) int64 {
	banM := make(map[int]bool)
	now := 0
	var ans int64 = 0
	for i := 0; i < n; i++ {
		for {
			now++
			if _, ok := banM[now]; ok {
				continue
			}
			break
		}
		banM[target-now] = true
		ans += int64(now)
	}
	return ans
}
func minOperations(nums []int, target int) int {
	have := make([]int, 35)
	need := make([]int, 35)
	map2 := make(map[int]int)
	tt := 1
	ans := 0
	for t := 0; t <= 30; t++ {
		map2[tt] = t
		tt *= 2
	}
	for _, v := range nums {
		have[map2[v]]++
	}
	now := 0
	for target > 0 {
		if target&1 == 1 {
			need[now]++
		}
		target /= 2
		now++
	}
	for i := 0; i < 34; i++ {
		if need[i] == 1 {
			if have[i] > 0 {
				have[i+1] += (have[i] - 1) / 2
				have[i] = 1
			} else {
				//借
				j := i + 1
				for ; j < 34; j++ {
					if have[j] > 0 {
						have[j]--
						have[i] = 2
						ans += j - i
						break
					}
					have[j] = 1
				}
				if j == 34 {
					return -1
				}
			}
		} else {
			if have[i] > 0 {
				have[i+1] += have[i] / 2
				have[i] = have[i] % 2
			}
		}
	}
	return ans
}
