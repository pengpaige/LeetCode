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
func swapPairs(head *ListNode) *ListNode {
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

