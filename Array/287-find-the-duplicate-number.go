func findDuplicateBF(nums []int) int {
    ans := 0
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[j] == nums[i] {
                ans = nums[i]
            }
        }
    }
    return ans 
}


// 参考这个 https://dwz1.cc/XtUOSsD
func findDuplicateBinarySearch(nums []int) int {
    left, right := 1, len(nums)-1
    for left < right {
        mid := (left+right)/2
        count := 0
        for _, n := range nums {
            if n <= mid {
                count++
            }
        }
        if count > mid {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left
}


// 二进制
func findDuplicate(nums []int) int {
    mask, ans := 0, 0
    for i := 0; i < 32; i++ {
        mask = 1 << i
        a, b := 0, 0
        for index := 0; index < len(nums); index++ {
            if (nums[index] & mask) > 0 {
                a++
            }
            if (index & mask) > 0 {
                b++
            }
        }
        if a > b {
            ans = ans | mask
        }
    }
    return ans
}
