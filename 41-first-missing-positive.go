// 不用想了，这类题断然是想不出来的，看答案都得看半天
// 缺少的正数在 1~len(nums)+1 之间这一点好理解
// 但是基于这一点再想办法找到缺失的最小正数就太 trick 了
// nums[nums[i]%n] += n 这句来说，其实在第一次循环结束之后就没有大于 n 的数字了
// 但是即使 append 了一个元素之后 nums 下标最大还是 n-1 也就是原来 nums 的长度
// 原作者的想法就是把 遍历到的元素出现的次数记录到以该元素为下标的 nums 位置上
// 具体就是出现多少次就加上几次 n，最后通过除以 n 来判断是否遍历到了该元素
// if nums[i] < 0 || nums[i] >= n {
//     nums[i] = 0
// }
// 上面需要排除 大于等于 n 的情况是因为原数长长度是 n-1，增加新元素之后长度 n 不在新的下标范围内
// 如果去掉取余的逻辑是否可行呢？
// 不行，第二个 for 的遍历过程中可能出现元素值大于 n， 不取余的话会越界
func firstMissingPositive(nums []int) int {
    nums = append(nums, 0)
    n := len(nums)
    for i := 0; i < n; i ++ {
        if nums[i] < 0 || nums[i] >= n {
            nums[i] = 0
        }
    }
    for i := 0; i < n; i ++ {
        nums[nums[i]%n] += n
    }
    for i := 0; i < n; i ++ {
        if nums[i]/n == 0 {
            return i
        }
    }
    return n
}
