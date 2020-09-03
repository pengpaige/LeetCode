/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 使用哨兵的方法
// 哨兵并不能降低整个算法的复杂度
// 但是可以统一第一次操作和后续操作
// 因为如果想方便的进行后续操作, 就需要一直保持一个 pre 用来记录上一组交换之后的第二个节点, 用来连接前后交换的节点
// 而第一次交换是没有这个节点的
// 使用哨兵的话, 就可以不对第一组交换的节点做特殊处理
func swapPairs(head *ListNode) *ListNode {
	shead := &ListNode{Next: head}
	s := shead
	for s.Next != nil && s.Next.Next != nil {
		a, b := s.Next, s.Next.Next
		s.Next, a.Next, b.Next = b, b.Next, a
		s = a
	}
	return shead.Next
}

func swapPairs__(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 如果 list 长度是奇数则无需反转最后一个 node
	ans, prev, curr := head, head, head.Next
	for curr != nil {
		if prev == head {
			ans = curr
		}
		next := curr.Next
		// prev.Next 暂时指向 next 防止出现死循环
		curr.Next, prev.Next = prev, next
		if next == nil || next.Next == nil {
			break
		}
		prev.Next, curr, prev = next.Next, next.Next, next
	}
	return ans
}

func swapPairs_(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 如果 list 长度是奇数则无需反转最后一个 node
	ans, prev, curr := head, head, head.Next
	for curr != nil {
		if prev == head {
			ans = curr
		}
		next := curr.Next
		curr.Next = prev
		prev.Next = next // 暂时指向 next 防止出现死循环
		if next == nil || next.Next == nil {
			break
		}
		prev.Next = next.Next
		curr = next.Next
		prev = next
	}
	return ans
}

// @lc code=end

