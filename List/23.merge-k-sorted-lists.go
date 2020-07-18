import "container/heap"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var fhead = &ListNode{}
	h := &IntHeap{}
	heap.Init(h)

	for i := range lists {
		if lists[i] == nil {
			continue
		}
		heap.Push(h, node{Val: lists[i].Val, ListIdx: i})
		lists[i] = lists[i].Next
	}

	curr := fhead
	for h.Len() > 0 {
		n := heap.Pop(h).(node)
		curr.Next = &ListNode{Val: n.Val}
		if lists[n.ListIdx] != nil {
			heap.Push(h, node{
				Val:     lists[n.ListIdx].Val,
				ListIdx: n.ListIdx})
			lists[n.ListIdx] = lists[n.ListIdx].Next
		}
		curr = curr.Next
	}
	return fhead.Next
}

type node struct {
	Val     int
	ListIdx int
}

type IntHeap []node

func (h IntHeap) Len() int { return len(h) }

// Less 递增表示构成小顶堆
func (h IntHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(node))
}

func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

// @lc code=end

