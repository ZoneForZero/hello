package main

import (
	"fmt"
	"math"
	"sort"
)

func gcd(a, b int) int {
	if b%a == 0 {
		return a
	}
	return gcd(b%a, a)
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
func getC(num, all int) int {
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
func getTuFromEdge(edges [][]int, numNode int) [][]int {
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
func main() {
	k := 1000000000 * 5
	ans := 0
	k
	// fmt.Println(getTuFromEdge([][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}}, 5))
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)
	objs := make([][2]int, n)
	for i := 0; i < n; i++ {
		objs[i] = [2]int{difficulty[i], profit[i]}
	}
	sort.Slice(objs, func(i, j int) bool {
		return objs[i][1] > objs[j][1]
	})
	sort.Slice(worker, func(i, j int) bool {
		return worker[i] < worker[j]
	})
	minDai := min(worker[m-1], objs[0][0])
	ans := 0
	maxR := m
	for _, obj := range objs {
		if obj[0] > minDai {
			continue
		}
		l, r := 0, maxR
		// 找到刚好worker[z] > obj[0]的点
		for l < r {
			z := (l + r) / 2
			if worker[z] >= obj[0] {
				r = z
			} else {
				l = z + 1
			}
		}
		ans += (maxR - r) * obj[1]
		maxR = r
		if r > 0 && worker[r-1] < obj[0] {
			minDai = worker[r-1]
		} else {
			minDai = obj[0]
		}
	}
	return ans
}

func main() {
	var num, sum, ans, mi, ma int64
	fmt.Scan(&num)
	a := make([]int64, num)
	for i := int64(0); i < num; i++ {
		fmt.Scan(&a[i])
		ans += a[i]
	}
	sum = 0
	for i := int64(0); i < num; i++ {
		sum += a[i]
		if sum > 0 {
			sum = 0
		} else if sum < mi {
			mi = sum
		}
	}
	sum = 0
	for i := int64(0); i < num; i++ {
		sum += a[i]
		if sum < 0 {
			sum = 0
		} else if sum > ma {
			ma = sum
		}
	}
	if ma+mi > ans {
		ans = ma
	} else {
		ans = ans - mi
	}
	fmt.Println(ans)
}
func maxSubarraySumCircular(nums []int) int {
	var num, sum, ans, mi, ma int64
	fmt.Scan(&num)
	a := make([]int64, num)
	for i := int64(0); i < num; i++ {
		fmt.Scan(&a[i])
		ans += a[i]
	}
	sum = 0
	for i := int64(0); i < num; i++ {
		sum += a[i]
		if sum > 0 {
			sum = 0
		} else if sum < mi {
			mi = sum
		}
	}
	sum = 0
	for i := int64(0); i < num; i++ {
		sum += a[i]
		if sum < 0 {
			sum = 0
		} else if sum > ma {
			ma = sum
		}
	}
	if ma+mi > ans {
		ans = ma
	} else {
		ans = ans - mi
	}
}
