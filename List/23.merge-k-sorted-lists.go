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
	// container/heap 会自动帮你做 heapify
	// 每次从小顶堆 pop 之后要从原来的 list 中再取出一个 node
	// 如果原来的 list 为空就继续 pop
	// 这也就是为什么不直接存 ListNode.Val 到 heap 里
	// 为了保存当前从堆 pop 出来的值来自哪个 list 专门声明了 node 类型
	// 另外，可能你会对为什么一定要从 pop 出来的值原来的 list 再取一个 node
	// 原因在于，本地的 merge 逻辑就是随时都在堆里放当前所有 list 的最小值
	// 每次都从原 list 取一个值放进堆里，这个值被 push 进去之后可能不会成为堆顶
	// 但是成为一旦原 list 的下一个值可能成为堆顶或者可能在其被 push 进堆之前成为堆顶
	// 就会造成最终 merge 排序的乱序
	// 所以最稳健和简单的做法就是每次都从原 list 找下一个值 push 到堆里
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

