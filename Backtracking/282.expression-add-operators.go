/*
 * @lc app=leetcode id=282 lang=golang
 *
 * [282] Expression Add Operators
 */

// @lc code=start
import "strconv"

func addOperators(num string, target int) []string {
	ans := make([]string, 0)
	if len(num) == 0 {
		return ans
	}

	bt282("", num, target, 0, 0, 0, &ans)
	return ans
}

// @prevDelta 上一次操作（加减乘）造成的 currSum 变化值，方便遇到乘法时根据计算顺序回退计算结果
func bt282(sub, nums string, target, start, prevDelta, currSum int, ans *[]string) {
	if start == len(nums) && target == currSum {
		*ans = append(*ans, sub)
		return
	}
	for i := start; i < len(nums); i++ {
		currStr := string(nums[start : i+1])
		currNum, _ := strconv.Atoi(currStr)
		// 下面两个 corner case 顺序不能调整
		// 因为可能 nums 开头就是 0 甚至是连续的 0 需要排除后面生成 0 开头数字的情况
		if nums[start] == '0' && i > start {
			break
		}
		if start == 0 {
			bt282(sub+currStr, nums, target, i+1, currNum, currSum+currNum, ans)
			continue
		}
		bt282(sub+"+"+currStr, nums, target, i+1, currNum, currSum+currNum, ans)
		bt282(sub+"-"+currStr, nums, target, i+1, -1*currNum, currSum-currNum, ans)
		bt282(sub+"*"+currStr, nums, target, i+1, prevDelta*currNum, currSum-prevDelta+prevDelta*currNum, ans)
	}
}

// @lc code=end

