/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
// DP 就是这么枯燥且无味，毫无优化就 faster than 71.14%
func lengthOfLIS(nums []int) int {
	// dp[i] 表示截止到 nums[i] 时 LIS 的长度
    // 需要遍历一遍 dp 找到最大值
	dp := make([]int, len(nums))
	// base case 看起来很像废话，但 LIS 的长度确实最少是 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}    
    var ans int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
        ans = max(ans, dp[i])
	}
	return ans
}

//  这是一次 DFS 的超低空飞行，危险刺激但是能 AC
func lengthOfLIS_DFS(nums []int) int {
	var dfs func(s, prev, count int) int
	// 直接递归会超时，必须加记忆
	// key 表示 nums 字符起始位置
	// value 表示从 key 代表的位置开始到 nums 结尾 LIS 的长度增量
	mem := make(map[int]int, 0)
	// s 正在处理的子串起始下标
	// prev 递归树中上一个层级已构成的上升序列最大值（最后一个元素值）
	// count nums[s] 加入到递归分支的 LIS 之后 LIS 的长度
	// 这样之所以可以记忆中间结果是因为
	// 不同的递归求解 LIS 的过程可能会经过相同的 nums[s] 
	// 只要当前到达的 nums[s] 相同的话，后续能将上升序列长度增加多少
	// 完全取取决于 nums[s] 之后的子数组构成
	// 这种记忆方式已经无限接近 DP 了
	dfs = func(s, prev, count int) int {
        if _, ok := mem[s]; ok {
            return count+mem[s]
        }
		var ret int
        flag := false
        for i := s; i < len(nums); i++ {
			if nums[i] > prev {
                flag = true
				ret = max(dfs(i+1, nums[i], count+1), ret)
			}
		}
        if !flag {
            ret = count-1
        }
        mem[s] = ret - count
		return ret
	}
	return dfs(0, -1<<31, 1)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 双指针遍历数组的方式无法应对短期出现上升趋势但不是「最长」上升序列的情况
// 比如第 20/24 测试例 [10,9,2,5,3,4]
// 在 2, 5 这里出现上升趋势，但实际上 LIS 应该是 2, 3, 4
// 如果想穷举所有可能必须有「回退」指针的能力，可能回溯是唯一的朴素方法
func lengthOfLIS_WA_TwoPonits(nums []int) int {
	sz := len(nums)
	if sz == 0 {
		return 0
	}
	var ans int
    for i := 0; i < sz; i++ {
		j, tmp, curr := i+1, nums[i], 1
		for j < sz {
			if nums[j] > tmp {
				tmp = nums[j]
				j++
				curr++
			} else {
				j++
			}
		}
		if ans < curr {
			ans = curr
		}
	}
	return ans
}
// @lc code=end

