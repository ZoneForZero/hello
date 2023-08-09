package lc

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}
	min1, min2, minK := grid[0][0], grid[0][1], 0
	if min2 < min1 {
		min1, min2, minK = min2, min1, 1
	}
	for i := 2; i < n; i++ {
		nv := grid[0][i]
		if nv < min2 {
			if nv < min1 {
				min1, min2, minK = nv, min1, i
			} else {
				min2 = nv
			}
		}
	}
	for i := 1; i < n; i++ {
		nmin1, nmin2, nmink := 0, 0, 0
		{
			// 0
			nv := grid[i][0]
			if 0 == minK {
				nv += min2
			} else {
				nv += min1
			}
			nmin1 = nv
			// 1
			nv = grid[i][1]
			if 1 == minK {
				nv += min2
			} else {
				nv += min1
			}
			if nv < nmin1 {
				nmin1, nmin2, nmink = nv, nmin1, 1
			} else {
				nmin2, nmink = nv, 0
			}
		}
		for j := 2; j < n; j++ {
			nv := grid[i][j]
			if minK == j {
				nv += min2
			} else {
				nv += min1
			}
			if nv < nmin2 {
				if nv < nmin1 {
					nmin1, nmin2, nmink = nv, nmin1, j
				} else {
					nmin2 = nv
				}
			}
		}
		min1, min2, minK = nmin1, nmin2, nmink
	}
	return min1
}
func init() {

}
