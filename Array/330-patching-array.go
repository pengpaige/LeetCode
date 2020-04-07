// https://dwz1.cc/VzNvPD2
func minPatches(nums []int, n int) int {
    ans, miss, i := 0, 1, 0
    for miss <= n {
        if i < len(nums) && nums[i] <= miss {
            miss += nums[i]
            i++
        } else {
            miss *= 2
            ans++
        }
    }
    return ans
}
