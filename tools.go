package main

import (
	"math"
	"sort"
)

// 最大公约数
func gcd(a, b int) int {
	if b%a == 0 {
		return a
	}
	return gcd(b%a, a)
}

// 最小公倍数
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 取最大值
func maxInt(ints ...int) (rt int) {
	rt = ints[0]
	for _, v := range ints {
		if v > rt {
			rt = v
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 取最小值
func minInt(ints ...int) (rt int) {
	rt = ints[0]
	for _, v := range ints {
		if v < rt {
			rt = v
		}
	}
	return
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 快速幂
func quickMi(base, mi, modVal int) int {
	ans := 1
	tempbase := base
	for mi > 0 {
		if mi&1 == 1 {
			ans = (ans * tempbase) % modVal
		}
		mi /= 2
		tempbase = (tempbase * tempbase) % modVal
	}
	return ans
}

// 质数判断
func isPrime(val int) bool {
	if val < 2 {
		return false
	}
	if val == 2 || val == 3 {
		return true
	}
	if val%2 == 0 {
		return false
	}
	maxV := int(math.Sqrt(float64(val)))
	for i := 5; i <= maxV; i += 2 {
		if val%i == 0 {
			return false
		}
	}
	return true
}

// 筛法 rangeMax内的质数数组
func getPrimes(rangeMax int) []int {
	if rangeMax >= 10000000 {
		return sieveOfEuler(rangeMax)
	}
	return sieveOfEratosthenes(rangeMax)
}

// 质数筛法
func sieveOfEratosthenes(n int) []int {
	notPrimes := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !notPrimes[i] {
			for j := i * i; j <= n; j += i {
				notPrimes[j] = true
			}
		}
	}
	result := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !notPrimes[i] {
			result = append(result, i)
		}
	}
	return result
}

// 欧拉筛法 数据大yu1e7的时候会更快
func sieveOfEuler(n int) []int {
	primes := make([]bool, n+1)
	result := make([]int, 0)

	for i := 2; i <= n; i++ {
		if !primes[i] {
			result = append(result, i)
		}

		for _, prime := range result {
			if i*prime > n {
				break
			}
			primes[i*prime] = true

			if i%prime == 0 {
				break
			}
		}
	}

	return result
}

// 前缀和
func getPreSums(arr *[]int) []int {
	lenA := len(*arr)
	preSum := make([]int, 1, lenA+1)
	sum := 0
	for _, v := range *arr {
		sum += v
		preSum = append(preSum, sum)
	}
	return preSum
}

// 获取少于n的完全平方数
func getPerfectSquares(n int) []int {
	ans := []int{}
	i := 1
	ji := 1
	for {
		ji = i * i
		if ji > n {
			return ans
		}
		ans = append(ans, ji)
		i++
	}
}

// 无限背包求数量
func allToNum(sum int, packs []int) int {
	dp := make([]int, sum+1)
	dp[0] = 0
	maxV := 9999999
	for i := sum; i > 0; i-- {
		dp[i] = maxV
	}
	sort.Slice(packs, func(i, j int) bool {
		return packs[i] < packs[j]
	})
	for i := len(packs) - 1; i >= 0; i-- {
		for j := packs[i]; j <= sum; j++ {
			if dp[j-packs[i]]+1 < dp[j] {
				dp[j] = dp[j-packs[i]] + 1
			}
		}
	}
	return dp[sum]
}

// 组合
func getC(num, all int) int {
	if all-num > num {
		return getC(all-num, num)
	}
	cheng := all
	chu := 1
	ans := 1
	for num > 0 {
		num--
		ans = ans * cheng / chu
		chu++
		cheng--
	}
	return ans
}

// 边构造无向图
func lineToTu(edges [][]int, numNode int) [][]int {
	ans := make([][]int, numNode)
	for i := 0; i < numNode; i++ {
		ans[i] = make([]int, numNode)
	}
	lenE := len(edges)
	tuFlags := make([]bool, numNode)
	for i := 0; i < lenE; i++ {
		node1, node2 := edges[i][0], edges[i][1]
		if tuFlags[node1] {
			for j := numNode - 1; j >= 0; j-- {
				if ans[node1][j] != 0 {
					newVal := ans[node1][j] + 1
					ans[j][node2], ans[node2][j] = newVal, newVal
				}
			}
		}
		if tuFlags[node2] {
			for j := numNode - 1; j >= 0; j-- {
				if ans[node1][j] == 0 && ans[node2][j] != 0 {
					newVal := ans[node2][j] + 1
					ans[j][node1], ans[node1][j] = newVal, newVal
				}
			}
		}
		ans[node1][node2], ans[node2][node1] = 1, 1
		tuFlags[node1], tuFlags[node2] = true, true
	}
	return ans
}

// 两int的绝对差
func seq(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 坐标点的曼哈顿距离计算
func pointMHD(point1, point2 []int) int {
	return seq(point1[0], point2[0]) + seq(point1[1], point2[1])
}

// 坐标点数组生成 点距离
func pointToTu(points [][]int) [][]int {
	lenP := len(points)
	tu := make([][]int, lenP)
	for i := 0; i < lenP; i++ {
		tu[i] = make([]int, lenP)
	}
	for i := 0; i < lenP; i++ {
		for j := i + 1; j < lenP; j++ {
			v := pointMHD(points[i], points[j])
			tu[i][j] = v
			tu[j][i] = v
		}
	}
	return tu
}

// prim最小生成树
func prim(tu [][]int) int {
	// 图与点的最短距离 初始化默认为下标0的距离关系
	minVs := tu[0]
	nodeNum := len(tu)
	nowNum := 1
	maxV := 1000000 * 3
	ans := 0
	for nowNum < nodeNum {
		nowNum++
		nextV := maxV
		nextIndex := 0
		for i := 0; i < nodeNum; i++ {
			if minVs[i] == 0 {
				continue
			}
			if minVs[i] < nextV {
				nextV = minVs[i]
				nextIndex = i
			}
		}
		// nextIndex -> tu
		ans += minVs[nextIndex]
		for i := 0; i < nodeNum; i++ {
			if tu[nextIndex][i] < minVs[i] {
				minVs[i] = tu[nextIndex][i]
			}
		}
	}
	return ans
}

// 小根堆
type gz struct {
	X, Y, V int
}
type GzHeap []gz

func (gzheap GzHeap) Len() int            { return len(gzheap) }
func (gzheap GzHeap) Less(i, j int) bool  { return gzheap[i].V < gzheap[j].V }
func (gzheap GzHeap) Swap(i, j int)       { gzheap[i], gzheap[j] = gzheap[j], gzheap[i] }
func (gzheap *GzHeap) Push(x interface{}) { *gzheap = append(*gzheap, x.(gz)) }

// 返回最后一个值，删除最后一个值
func (gzheap *GzHeap) Pop() interface{} {
	old := *gzheap
	n := len(old) - 1
	x := old[n]
	*gzheap = old[:n]
	return x
}

// minHeap := &GzHeap{gz{0,0,grid[0][0]}}
// heap.Init(minHeap)
// tempDian := heap.Pop(minHeap).(gz)
// heap.Push(minHeap,gz{nowX,nowY,nowV})

// func search(){
//         index := sort.Search(len(tt),func(z int) bool {
//             // 返回符合条件的最小index
//             return tt[z] > 1
//         })
// }
