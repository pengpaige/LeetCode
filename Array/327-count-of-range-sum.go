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
