/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p, len := head, 0
	for p != nil {
		len++
		p = p.Next
	}

	var pprev, pcurr *ListNode
	var qprev, qcurr *ListNode
	for i := 0; i < len; i++ {
		pprev, pcurr = nil, head
		qprev, qcurr = nil, head
		for j := 0; j <= i; j++ {
			qprev, qcurr = qcurr, qcurr.Next
		}
		if qcurr == nil {
			break
		}
		// for j := i+1; j < len: j++ {}
		for j := 0; j <= i; j++ {
			if pcurr.Val <= qcurr.Val {
				pprev, pcurr = pcurr, pcurr.Next
			} else {
				qprev.Next = qcurr.Next
				if pprev == nil {
					qcurr.Next = pcurr
					head = qcurr
				} else {
					pprev.Next = qcurr
					qcurr.Next = pcurr
				}
				break
			}
		}
	}

	return head
}

// @lc code=end