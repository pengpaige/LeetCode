// currPos 表示当前能跳到的最右位置
// 可以出现一次或者几次节点没有讲 currPos 右移的情况
// 但是一旦 i > currPos 就代表数组不满足要求，序号为 i 的节点不可达
func canJump1(nums []int) bool {
    var currPos int
    for i := 0; i < len(nums); i++ {
        if i > currPos {
            return false
        }
        currPos = max(i+nums[i], currPos)
    }
    return true
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


// lastPos 表示能跳到数组末尾的最右位置
func canJump(nums []int) bool {
    lastPos := len(nums)-1
    for i := len(nums)-2; i >= 0; i-- {
        // 可以出现某一次或者几次遇到有节点不能把 lastPos 左移
        // 但必须要有其他节点可以使 lastPos 可达数组最左边
        if i+nums[i] >= lastPos {
            lastPos = i
        }
    }
    return lastPos == 0
}
