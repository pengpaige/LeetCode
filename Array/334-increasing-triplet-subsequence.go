// brute force
func increasingTriplet1(nums []int) bool {
    for i := 0; i < len(nums)-1; i++ {
        temp := make([]int, 0)
        for j := i+1; j < len(nums); j++ {
            if nums[j] > nums[i] {
                temp = append(temp, nums[j])
            }
        }
        if len(temp) >= 3 {
            return true
        }
    }
    return false
}


func increasingTriplet(nums []int) bool {
    n := len(nums)
    if n < 3 {
        return false
    }
    small, mid := 1<<31-1, 1<<31-1
    for i := 0; i < n; i++ {
        // 注意这里要包含等于 small 否则错误地替换 mid
        if nums[i] <= small {
            small = nums[i]
        } else if nums[i] < mid {
            mid = nums[i]
        } else if nums[i] > mid {
            return true   
        }
    }
    return false
}
