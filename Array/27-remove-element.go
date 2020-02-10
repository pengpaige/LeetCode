// 这里使用的是头尾双指针的方式
// 使用快慢指针也可以
func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    l, r := 0, len(nums)-1
    for l < r {
        if nums[l] != val {
            l++
            continue
        }
        if nums[r] == val {
            r--
            continue
        }
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
    if nums[l] != val {
        return l+1
    }else{
        return l
    }
}