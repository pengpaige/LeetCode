/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var mergeSort func(head *ListNode) *ListNode
	mergeSort = func(node *ListNode) *ListNode {
		if node.Next == nil {
			return node
		}

		pprev, p, px2 := node, node, node
		for px2 != nil {
			pprev, p, px2 = p, p.Next, px2.Next
			if px2 != nil {
				px2 = px2.Next
			} else {
				break
			}
		}
		// 这里要注意必须把当前需要处理的链表段分开，否则会无限递归
		pprev.Next = nil

		left, right := mergeSort(node), mergeSort(p)
		var newHead, newp, min *ListNode
		for left != nil || right != nil {
			if left != nil && right != nil {
				if left.Val < right.Val {
					min = left
					left = left.Next
				} else {
					min = right
					right = right.Next
				}
			} else if left != nil {
				min = left
				left = left.Next
			} else {
				min = right
				right = right.Next
			}
			if newHead == nil {
				newHead, newp = min, min
			} else {
				newp.Next = min
				newp = newp.Next
			}
		}
		return newHead
	}
	return mergeSort(head)
}

// @lc code=end

