// 空间复杂度 O(k) 的方法，额外创建 slice 存储搬移到头部的元素
func rotate1(nums []int, k int)  {
    // k 可能超过 nums 的长度
    k=k%len(nums)
    s1 := make([]int, k)
    for i := 0; i < k; i++ {
        s1[i] = nums[len(nums)-k+i]
    }
    // 前面的数字向后搬移的话，得从尾部开始，防止尾部的数字还没被搬移就已经被前面先搬移的数字覆盖
    for j := len(nums)-k-1; j >= 0; j-- {
        nums[j+k]=nums[j]
    }
    for i := 0; i < k; i++ {
        nums[i] = s1[i]
    }
}

// 最佳方法，通过翻转实现原地旋转
func rotate(nums []int, k int)  {
    n := len(nums)
    k %= n
    reverse(nums, 0, n-1)
    reverse(nums, 0, k-1)
    reverse(nums, k, n-1)
}


func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}

