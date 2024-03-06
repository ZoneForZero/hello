package lc

import (
	"fmt"
	"sort"
)

// 给你一个整数 n 表示数轴上的房屋数量，编号从 0 到 n - 1 。
//
// 另给你一个二维整数数组 offers ，其中 offers[i] = [starti, endi, goldi] 表示第 i 个买家想要以 goldi 枚金币的价格购买从 starti 到 endi 的所有房屋。
//
// 作为一名销售，你需要有策略地选择并销售房屋使自己的收入最大化。
//
// 返回你可以赚取的金币的最大数目。
//
// 注意 同一所房屋不能卖给不同的买家，并且允许保留一些房屋不进行出售。
func maximizeTheProfit(n int, offers [][]int) int {
	sort.Slice(offers, func(i, j int) bool {
		return offers[i][1] < offers[j][1]
	})
	lenO := len(offers)
	ans := 0
	dp := make([]int, lenO)
	for i := 0; i < lenO; i++ {
		tempSum := offers[i][2]
		start, j := 0, i
		for start < j {
			mid := (start + j) / 2
			if offers[mid][1] < offers[i][0] {
				start = mid + 1
			} else {
				j = mid
			}
		}
		j--
		// 记录最大起点
		maxQi := 0
		for ; j >= 0; j-- {
			if offers[j][1] < maxQi {
				break
			}
			if offers[j][0] > maxQi {
				maxQi = offers[j][0]
			}
			if dp[j]+offers[i][2] > tempSum {
				tempSum = dp[j] + offers[i][2]
			}
		}
		//for j := 0; j < i; j++ {
		//	if offers[j][1] >= offers[i][0] {
		//		break
		//	}
		//	if dp[j]+offers[i][2] > tempSum {
		//		tempSum = dp[j] + offers[i][2]
		//	}
		//}
		if tempSum > ans {
			ans = tempSum
		}
		dp[i] = tempSum
	}
	return ans
}
func init() {
	fmt.Println("ok")

}
func longestEqualSubarray(nums []int, k int) int {
	ma := make(map[int][]int)
	chaMap := make(map[int][]int)
	lenN := len(nums)
	for i := 0; i < lenN; i++ {
		val := nums[i]
		if _, ok := ma[val]; !ok {
			ma[val] = []int{i}
		} else {
			ma[val] = append(ma[val], i)
		}
	}
	for k, indexArr := range ma {
		if len(indexArr) < 2 {
			continue
		}
		chaMap[k] = []int{}
		lenI := len(indexArr)
		for i := 1; i < lenI; i++ {
			chaMap[k] = append(chaMap[k], indexArr[i]-indexArr[i-1]-1)
		}
	}
	ans := 1
	for _, indexArr := range chaMap {
		if len(indexArr) <= ans {
			continue
		}
		lenA := len(indexArr)
		s := 0
		e := 0
		dj := 0
		tAns := 0
		for s < lenA {
			// 直到不行
			for dj <= k && e < lenA {
				dj += indexArr[e]
				e++
			}
			if e-s+1 > tAns {
				tAns = e - s + 1
			}
			if e == lenA {
				break
			}
			// 锁
			for dj > k {
				dj -= indexArr[s]
				s++
			}
		}
		if tAns > ans {
			ans = tAns
		}
	}
	return ans
}
