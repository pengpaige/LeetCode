/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Brute Force
func reorderList(head *ListNode) {
	p := head
	for p != nil {
		var prev *ListNode
		curr := p
		for curr.Next != nil {
			prev = curr
			curr = curr.Next
		}
		if curr == p || curr == p.Next {
			return
		}
		p.Next, curr.Next = curr, p.Next
		prev.Next = nil
		p = p.Next.Next
	}
	return
}

// @lc code=end

