/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	p := head
	l := 0
	for p != nil {
		p = p.Next
		l++
	}
	if l < 2 {
		return head
	}

	kk := k%l
	// 旋转长度是原长度的整数倍就直接返回原来的 head
	if kk == 0 {
		return head
	}
	pprev, pcurr := head, head.Next
    // 注意移动的距离，这个自己画图推导一下就清楚了
    for i := l-(kk+1); i > 0; i-- {
		pprev, pcurr = pcurr, pcurr.Next
	}
	// 从旋转位置段开链表
	pprev.Next = nil
    
	newHead := pcurr
	for pcurr.Next != nil {
		pcurr = pcurr.Next
	}
	pcurr.Next = head

	return newHead
}
// @lc code=end

