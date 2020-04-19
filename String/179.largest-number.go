/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
import(
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	allZeroFlag := true
	for _, num := range nums {
		if num != 0 {
			allZeroFlag = false
		}
	}
	if allZeroFlag {
		return "0"
	}

    sort.Slice(nums, func(i, j int) bool {
		si := strconv.Itoa(nums[i])
		sj := strconv.Itoa(nums[j])
		
		for {
            d := 0
			for ; d < min(len(si), len(sj)); d++ {
				if si[d] == sj[d] {
					continue
				} else {
					return si[d] > sj[d]
				}
			}
			if len(si) > len(sj) {
				si = si[d:]
			} else if len(si) < len(sj) {
				sj = sj[d:]
			} else {
				break
			}
		}

		return false
	})
	
	var ans string
	for _, num := range nums {
		ans += strconv.Itoa(num)
	}
	return ans
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

