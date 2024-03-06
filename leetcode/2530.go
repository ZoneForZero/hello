package lc

import (
	"container/heap"
)

type GzHeap []int

func (gzheap GzHeap) Len() int            { return len(gzheap) }
func (gzheap GzHeap) Less(i, j int) bool  { return gzheap[i] > gzheap[j] }
func (gzheap GzHeap) Swap(i, j int)       { gzheap[i], gzheap[j] = gzheap[j], gzheap[i] }
func (gzheap *GzHeap) Push(x interface{}) { *gzheap = append(*gzheap, x.(int)) }

// 返回最后一个值，删除最后一个值
func (gzheap *GzHeap) Pop() interface{} {
	old := *gzheap
	n := len(old) - 1
	x := old[n]
	*gzheap = old[:n]
	return x
}

func maxKelements(nums []int, k int) int64 {
	arr := &GzHeap{}
	*arr = nums
	var ans int64 = 0

	heap.Init(arr)
	for k > 0 {
		k--
		maxV := heap.Pop(arr).(int)
		ans += int64(maxV)
		if maxV%3 == 0 {
			heap.Push(arr, maxV/3)
		} else {
			heap.Push(arr, 1+(maxV/3))
		}

	}
	// heap.Push(minHeap, gz{nowX, nowY, nowV})
	return ans
}
