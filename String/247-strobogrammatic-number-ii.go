/*
这是付费，题目是从网上找到的。

A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

Find all strobogrammatic numbers that are of length = n.

Example:

Input:  n = 2
Output: ["11","69","88","96"]

下面是一些例子

n = 0:   none
n = 1:   0, 1, 8
n = 2:   11, 69, 88, 96
n = 3:   101, 609, 808, 906, 111, 619, 818, 916, 181, 689, 888, 986
n = 4:   1001, 6009, 8008, 9006, 1111, 6119, 8118, 9116, 1691, 6699, 8698, 9696, 1881, 6889, 8888, 9886, 1961, 6969, 8968, 9966

注意观察 n=0 和 n=2，可以发现后者是在前者的基础上，每个数字的左右增加了 [1 1], [6 9], [8 8], [9 6]，看 n=1 和 n=3 更加明显，在0的左右增加 [1 1]，变成了 101, 在0的左右增加 [6 9]，变成了 609, 在0的左右增加 [8 8]，变成了 808, 在0的左右增加 [9 6]，变成了 906, 然后在分别在1和8的左右两边加那四组数，实际上是从 m=0 层开始，一层一层往上加的，需要注意的是当加到了n层的时候，左右两边不能加 [0 0]，因为0不能出现在两位数及多位数的开头，在中间递归的过程中，需要有在数字左右两边各加上0的那种情况，

这里需要注意的一点是，上面说的一层一层叠加的思路，在得到最终的结果之前都是中间值，不是某一层真正的序列。
这一点从代码中对于递归过程中是否在两侧加 0 的处理就可以看出来。
举个例子，上面 n=4 时，第一个数字 1001 是通过对 00 这两个数两侧加 1 得到的，但是实际上 n=2 的时候并没有 00.package String

代码如下，每次n-2，一层一层的迭代，到最后一层，n = 0 或者 n = 1 停止，backtracking加对称数字。
n = 0时得到空串, n = 1时，得到 0 1 8，其它层每次加能组成对称数字对的一种组合。
*/

func findStrobogrammatic(int n) []string {
	return helper(n, n)
}

// cl: current level
// tl: total levels
func helper(cl, tl int) []string {
	if cl == 0 {
		return []string{ "" }
	} else if cl == 1 {
		return []string{ "0", "1", "8" }
	}
	var ret []string
	for _, num := range helper(cl-2, tl) {
		if cl != tl {
			ret = append("0"+num+"0")
		}
		ret = append("1"+num+"1")
		ret = append("8"+num+"8")
		ret = append("6"+num+"9")
		ret = append("9"+num+"6")
	}
	return ret
}