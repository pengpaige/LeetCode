// 从这里抄的，就不写注释了，直接看原文吧
// http://dwz1.cc/6KR3zJS
func jump(nums []int) int {
    if len(nums) <=1 {
        return 0
    }
    currMinStep, start := make([]int, len(nums)), make([]int, len(nums))
    // 因为一定可达末尾，所以也一定可达第 1 个点，且只需要一步
    currMinStep[1] = 1
    // 最少步数为 1 的起始位置也自然就是第 1 点
    start[1] = 1
    for i := 2; i < len(nums); i++ {
        currMinStep[i] = currMinStep[i-1]+1
        // start 先更新成 i，如果不是 i 的话下次外层循环会重新更新 start
        start[currMinStep[i]] = i
        for j := start[currMinStep[i-1]-1]; j < start[currMinStep[i-1]]; j++ {
            if j+nums[j] >= i {
                currMinStep[i] = currMinStep[i-1]
                break
            }
        }
    }
    return currMinStep[len(nums)-1]
}
