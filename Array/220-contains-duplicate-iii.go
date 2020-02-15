/**
 * 与前面 i 和 ii 两个题没什么太大区别
 * 只是不能像前两道题那样直接用 map 辅助判断是否出现过(应为是判断是否有满足条件的相等元素)】
 * 因为相等的话直接用 map 的映射就能判断，现在是判断是否在一个范围内，map 没办法直接存储范围了
 */

// O(n^2) 的方法
func containsNearbyAlmostDuplicateS1(nums []int, k int, t int) bool {
    for i := 0; i < len(nums); i++ {
        for j := max(i-k, 0); j < i; j++ {
            if abs(nums[i]-nums[j]) <= t {
                return true
            }
        }
    }
    return false
}

// 如果想减少嵌套循环造成的耗时，就得想办法像上两题那样想办法做一个映射
// 但是现在需要映射的不是数组元素是否存在而是 从数组元素到该元素的 +t 和 -t 的范围
// 首先滑动窗口还是需要的，其次考虑到映射印个范围，用桶的方式来做映射
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    // 测试例里竟然还有 t<0 ？
    if k == 0 || t < 0 {
        return false
    }
    m := make(map[int]int, k)
    w := t+1
    for i := 0; i < len(nums); i++ {
        b := getBucketId(nums[i], w)
        if _, ok := m[b]; ok {
            return true
        }
        if n, ok := m[b+1]; ok && abs(nums[i]-n) <= t {
            return true
        }
        if n, ok := m[b-1]; ok && abs(nums[i]-n) <= t {
            return true
        }
        if len(m) == k {
            delete(m, getBucketId(nums[i-k], w))
        }
        m[b] = nums[i]
    }
    return false
}

// 每个桶的范围是  [0, t] [t+1， 2t+1] [2t+2, 3t+2] ...
// 这个桶的范围，或者说叫宽度，对于这题来说是必须的选择，可以改变每个桶的其实位置但桶的宽度不能改
// 这样就能保证在一个桶内的数字直接满足差不大于 t，且相邻的桶内数字的差可能不大于 t
// 不相邻的桶内数字的差一定大于 t
func getBucketId(x, w int) int {
    if x < 0 {
        // 这句需要好好理解下
        return (x+1)/w - 1
    }
    return x/w
}


func abs(i int) int {
    if i < 0 {
        return -1*i
    }
    return i
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
