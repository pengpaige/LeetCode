/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
import (
	"container/heap"
)
// 小根堆
func findKthLargest(nums []int, k int) int {
	var h *IntHeap
	h = (*IntHeap)(&[]int{})
	for _, n := range nums {
		if h.Len() < k {
			heap.Push(h, n)
		} else if (*h)[0] < n { // len(h) = k
			heap.Pop(h)
			heap.Push(h, n)
		}
	}
	// for i := k-1; i > 0; i-- {
	// 	heap.Pop(h)
	// }
	return (*h)[0]
}

func findKthLargest_cheating(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// Less 递增表示构成小顶堆
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    ret := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return ret
}
// @lc code=end

