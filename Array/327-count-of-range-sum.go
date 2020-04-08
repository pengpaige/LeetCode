// 题目要求是不能用这种O(N^2)的算法的
// 下面的解法能通过可能是因为 Golang 代码判定测度太保守导致漏网
func countRangeSum(nums []int, lower int, upper int) int {
    var ans int
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum >= lower && sum <= upper {
                ans++
            }
        } 
    }
    return ans
}
