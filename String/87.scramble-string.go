/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */

// @lc code=start
import "sort"

func isScramble(s1 string, s2 string) bool {
	// mem := make(map[string]bool, 0)
	// return isScrambleRecursionBugfix(s1, s2)
	// return isScrambleRecursionMemorization(s1, s2, mem)
	return isScrambleRecursionDP(s1, s2)
}


func isScrambleRecursionDP(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	sortedS1 := string(sortByteSlice([]byte(s1)))
	sortedS2 := string(sortByteSlice([]byte(s2)))
	if sortedS1 != sortedS2 {
	return false
	}

	l := len(s1)
	dp := make([][][]bool, l+1)
	for ll := 0; ll <= l; ll++ {
		dp[ll] = make([][]bool, l)
		for i := 0; i < l; i++ {
			dp[ll][i] = make([]bool, l)
			// for j := 0; j + len < l; j++ {
			// 	dp[len][i][j] = false
			// }
		}
	}
	dp[0][0][0] = false

	for length := 1; length <= l; length++ {
		for i := 0; i + length <= l; i++ {
			for j := 0; j + length <= l; j++ {
				if length == 1 {
					dp[length][i][j] = s1[i] == s2[j]
				} else {
					for q := 1; q < length; q++ {
						dp[length][i][j] = dp[q][i][j] && dp[length-q][i+q][j+q] ||
							dp[q][i][j+length-q] && dp[length-q][i+q][j]
						if dp[length][i][j] == true {
							break
						}
					}
				}
			}
		}
	}
	return dp[l][0][0]
}


func isScrambleRecursionMemorization(s1, s2 string, mem map[string]bool) bool {
	if _, ok := mem[s1+"@"+s2]; ok {
		return mem[s1+"@"+s2]
	}
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	sortedS1 := string(sortByteSlice([]byte(s1)))
	sortedS2 := string(sortByteSlice([]byte(s2)))
	if sortedS1 != sortedS2 {
		return false
	}
	l := len(s1)
	for i := 1; i < l; i++ {
		if isScrambleRecursionMemorization(s1[:i], s2[:i], mem) && isScrambleRecursionMemorization(s1[i:l], s2[i:l], mem) {
			mem[s1+"@"+s2] = true
			return true
		}
		if isScrambleRecursionMemorization(s1[:i], s2[l-i:l], mem) && isScrambleRecursionMemorization(s1[i:l], s2[:l-i], mem) {
			mem[s1+"@"+s2] = true
			return true
		}
	}
	return false
}


func isScrambleRecursionBugfix(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	sortedS1 := string(sortByteSlice([]byte(s1)))
	sortedS2 := string(sortByteSlice([]byte(s2)))
	if sortedS1 != sortedS2 {
	return false
	}
	l := len(s1)
	for i := 1; i < l; i++ {
		if isScrambleRecursionBugfix(s1[:i], s2[:i]) && isScrambleRecursionBugfix(s1[i:], s2[i:]) {
			return true
		}
		if isScrambleRecursionBugfix(s1[:i], s2[l-i:]) && isScrambleRecursionBugfix(s1[i:], s2[:l-i]) {
			return true
		}
	}
	return false
}


func isScrambleRecursion(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sortedS1 := string(sortByteSlice([]byte(s1)))
	sortedS2 := string(sortByteSlice([]byte(s2)))
	// 这里的判断不应该放在排序之后，比较两者是否相同应该放在排序之前，因为不是所有包含字符相同但是顺序不同的字符串都能按题目要求转换得到
	if sortedS1 != sortedS2 {
		return false
	}
    if len(s1) == 1 {
        return true
    }
	l := len(s1)
	// 下面回溯去寻找乱序数组是否相等的时候，应该用原 string 的顺序，而不是排序后的
	for i := 1; i < l; i++ {
		if isScrambleRecursion(sortedS1[:i], sortedS2[:i]) && isScrambleRecursion(sortedS1[i:l], sortedS2[i:l]) {
			return true
		}
		if isScrambleRecursion(sortedS1[:i], sortedS2[l-i:l]) && isScrambleRecursion(sortedS1[i:l], sortedS2[:l-i]) {
			return true
		}
	}
	return false
}


func sortByteSlice(bs []byte) []byte {
    sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	return bs
}

// @lc code=end

