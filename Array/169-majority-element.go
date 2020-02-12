// 
// func majorityElement(nums []int) int {
//     if len(nums) == 1 {
//         return nums[0]
//     }
//     curr, count := nums[0], 1
//     for i, n := range nums {
//         if i == 0 {
//             continue
//         }
//         if curr == n {
//             count++
//         } else {
//             count--
//         }
//         if count == 0 {
//             curr = n
//             count = 1
//         }
//     }
//     return curr
// }


// count == 0 这句要跟后面 count 加减分开，不能在一次循环里做
// 否则会多计算一次使 count 减为零的数字
func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    curr, count := 0, 0
    for _, n := range nums {
        if count == 0 {
            curr = n
            count = 1
        } else if curr == n {
            count++
        } else {
            count--
        }
    }
    return curr
}
