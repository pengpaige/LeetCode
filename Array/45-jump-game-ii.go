// lastRightMost 表示上一次跳到的最右位置
// nextRightMost 表示基于上一次的最右位置与上上次最右位置之间的点
// 可以跳到的最右位置
// 当 i == lastRightMost 时，说明已经达到本次跳跃可达的最优位置
// 需要将 lastRightMost 指向下一次跳跃的最右位置
// 另外，需要注意的是 len(nums)-1 这个位置不能被 i 取到
// 否则可能导致多计一次跳跃总数，这点可以在纸上自己模拟一遍
// 本方法比较特殊，很难想到，全当做发散思维吧
func jump(nums []int) int {
    ans, lastRightMost, nextRightMost := 0, 0, 0
    for i := 0; i < len(nums)-1; i++ {
        nextRightMost = max(i+nums[i], nextRightMost)
        if i == lastRightMost {
            ans++
            lastRightMost = nextRightMost
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
