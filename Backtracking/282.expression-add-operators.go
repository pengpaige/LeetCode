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

func bt282(sub, nums string, target, start, prevDelta, currSum int, ans *[]string) {
	if start == len(nums) && target == currSum {
		*ans = append(*ans, sub)
		return
	}
	for i := start; i < len(nums); i++ {
		currStr := string(nums[start : i+1])
		currNum, _ := strconv.Atoi(currStr)

		bt282(sub+"+"+currStr, nums, target, i+1, currNum, currSum+currNum, ans)
		bt282(sub+"-"+currStr, nums, target, i+1, -1*currNum, currSum-currNum, ans)
		bt282(sub+"*"+currStr, nums, target, i+1, prevDelta*currNum, currSum-prevDelta+prevDelta*currNum, ans)
	}
}

// @lc code=end

