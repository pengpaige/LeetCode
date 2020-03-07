// brute force
func maxArea0(height []int) int {
    ans := 0
    for i := 0; i < len(height); i++ {
        for j := i+1; j < len(height); j++ {
            ans = max(ans, min(height[i], height[j])*(j-i))
        }
    }
    return ans
}


/* 双指针 http://dwz1.cc/MBoJ2tn
开始时两段各一个指针，每次把高度更小的指针向内移动，并和当前的最大面积比较
为什么必须是移动高度较小的指针需要自己思考
*/
func maxArea(height []int) int {
    ans := 0
    for i, j := 0, len(height)-1; i < j; {
        ans = max(ans, min(height[i], height[j])*(j-i))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
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
