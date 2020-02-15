// 这里需要注意题目要求是最大下标差等于 k
// 下面的方法忽略的窗口从 0 增大到 k+1 的过程
// 遇到输入长度小于 k+1 的情况就会出错
func containsNearbyDuplicateR1(nums []int, k int) bool {
    if k > len(nums)-1 || k == 0 {
        return false
    }
    for i := k; i < len(nums); i++ {
        m := make(map[int]struct{})
        for j := i-k; j <= i; j++ {
            _, ok := m[nums[j]]
            if ok {
                return true
            }
            m[nums[j]] = struct{}{}
        }
    }
    return false
}


// 考虑了窗口增大的过程但是复杂度太高 O(n^2)
// 遇到输入非常长的测试例会超时
func containsNearbyDuplicateR2(nums []int, k int) bool {
    if k == 0 {
        return false
    }
    for i := 0; i < len(nums); i++ {
        m := make(map[int]struct{})
        for j := max(i-k, 0); j <= i; j++ {
            _, ok := m[nums[j]]
            if ok {
                return true
            }
            m[nums[j]] = struct{}{}
        }
    }
    return false
}


// 而且上面的 map 实际上是不需要的
// 在每个窗口里(包括窗口增长过程中)小于 K+1 长度的窗口，都是用下标为 i 和 i 之前的数字比较
// 这样每次窗口移动就只需重复上面和新的 i 之前的数字比较的操作
// 不需要再存储每个数字到 map 在讲所有当前窗口的数字交叉的进行是否存在的判定(这点需要思考)
func containsNearbyDuplicateR3(nums []int, k int) bool {
    if k == 0 {
        return false
    }
    for i := 0; i < len(nums); i++ {
        for j := max(i-k, 0); j < i; j++ {
            if nums[i] == nums[j] {
                return true
            }
        }
    }
    return false
}


// 其实 R2 已经很接近正确答案了，既然用了 map 就不需要再用嵌套遍历
// 而是只维持 map 的长度为 k(注意 map 的长度是 k 不是 k+1，逻辑类似与前面比较 i 和 i 前面的数字是否相同)
// map 长度达到 k+1 之后就删除一个最先加入的元素来维持窗口大小
// 因为当未发现重复的时候 map 长度保持小于等于 k+1，这样就通过 map 维护了一个 k+1 长的窗口
func containsNearbyDuplicate(nums []int, k int) bool {
    if k == 0 {
        return false
    }
    m := make(map[int]struct{}, k+1)
    for i := 0; i < len(nums); i++ {
        if _, ok := m[nums[i]]; ok {
            return true
        }
        if len(m) == k{
            delete(m, nums[i-k])
        }
        m[nums[i]] = struct{}{}
    }
    return false
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
