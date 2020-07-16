/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != nil && p2 != nil {
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		if p2.Next != nil {
			p2 = p2.Next.Next
		} else {
			return false
		}
	}
	return false
}

func hasCycle_Map(head *ListNode) bool {
	m := make(map[*ListNode]struct{}, 0)
	p := head
	for p != nil {
		if _, ok := m[p]; ok {
			return true
		}
		m[p] = struct{}{}
		p = p.Next
	}
	return false
}

// @lc code=end

