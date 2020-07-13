/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList_Brief(head *ListNode) *ListNode {
	var newHead, pprev, p *ListNode
	pprev, p = nil, head
	for p != nil {
		next := p.Next
		if next == nil {
			newHead = p
		}
		p.Next = pprev
		pprev = p
		p = next
	}
	return newHead
}

func reverseList(head *ListNode) *ListNode {
	var newHead, pprev, p *ListNode
	pprev, p = nil, head
	for p != nil {
		if p.Next == nil {
			newHead = p
		}
		p.Next, pprev, p = pprev, p, p.Next
	}
	return newHead
}

// @lc code=end

