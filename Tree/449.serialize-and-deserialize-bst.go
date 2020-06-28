import (
	"strconv"
	"strings"
)
/*
 * @lc app=leetcode id=449 lang=golang
 *
 * [449] Serialize and Deserialize BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ret string
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		c, _ := strconv.Itoa(node.Val)
		ret += c + " "
	}
	postOrder(root)
	ret = string(ret[:len(ret)-1])
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nums := this.StrToIntSlice(data) 
	var de func(lower, upper, val int) *TreeNode
	de = func(lower, upper int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		val := nums[len(nums)-1]
		if val < lower || val > upper {
			return nil
		}
		nums = nums[:len(nums)-1]
		node := &TreeNode{Val: val}
		node.Right = de(val, upper)
		node.Left = de(lower, val)
		return node
	}
	return de(-1<<31, 1<<31-1)
}

func (this *Codec) StrToIntSlice(str string) []int {
	var ret []int
	strs := strings.Split(str, " ")
	for _, s := range strs {
		ret = append(ret, strconv.Atoi(s))
	}
	return ret
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

