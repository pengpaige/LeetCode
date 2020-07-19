/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	over, p, p1, p2 := 0, head, l1, l2

	for p1 != nil || p2 != nil || over != 0 {
		var d, d1, d2 int
		if p1 != nil {
			d1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			d2 = p2.Val
			p2 = p2.Next
		}
		d = d1 + d2 + over
		if d >= 10 {
			d, over = d%10, d/10
		} else {
			over = 0
		}
		// 这里不对 head 做特殊处理，但是最后 return 要返回 head.Next
		node := &ListNode{
			Val:  d,
			Next: nil,
		}
		p.Next = node
		p = node
	}
	return head.Next
}

func addTwoNumbers_(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	over, p, p1, p2 := 0, head, l1, l2
	var pprev *ListNode
	d, d1, d2 := p1.Val+p2.Val, p1.Val, p2.Val
	p1, p2 = p1.Next, p2.Next
	if d >= 10 {
		d, over = d%10, d/10
	}
	p.Val = d
	pprev = head

	for p1 != nil || p2 != nil || over != 0 {
		d, d1, d2 = 0, 0, 0
		if p1 != nil {
			d1 = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			d2 = p2.Val
			p2 = p2.Next
		}
		d = d1 + d2 + over
		if d >= 10 {
			d, over = d%10, d/10
		} else {
			over = 0
		}
		p = &ListNode{
			Val:  d,
			Next: nil,
		}
		pprev.Next = p
		pprev = p
	}
	return head
}

// @lc code=end

