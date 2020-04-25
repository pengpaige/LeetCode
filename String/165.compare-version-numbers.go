/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	vs1, vs2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i := 0; i < Min(len(vs1), len(vs2)); i++ {
		subVer1, _ := strconv.Atoi(vs1[i])
		subVer2, _ := strconv.Atoi(vs2[i])
		if subVer1 == subVer2 {
			continue
		}
		if subVer1 > subVer2 {
			return 1
		} else {
			return -1
		}
	}
	if len(vs1) == len(vs2) {
		return 0
	} else if len(vs1) > len(vs2) {
		for _, subVer := range vs1[len(vs2):] {
			if subVerInt, _ := strconv.Atoi(subVer); subVerInt != 0 {
				return 1
			}
		}
		return 0
	} else {
		for _, subVer := range vs2[len(vs1):] {
			if subVerInt, _ := strconv.Atoi(subVer); subVerInt != 0 {
				return -1
			}
		}
		return 0
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

