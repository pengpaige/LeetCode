// 把不重复的数字依次放到数组最前面
// i 代表 nums 数组上一次循环被覆盖的位置
// 具体逻辑的话自己模拟一个循环就明白了
func removeDuplicates(nums []int) int {
    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[i] {
            nums[i+1] = nums[j]
            i++
        }
    }
    return i+1
}