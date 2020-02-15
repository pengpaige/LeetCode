func containsDuplicate(nums []int) bool {
    var ans bool
    m := make(map[int]struct{}, len(nums))
    for _, n := range nums {
        if _, ok := m[n]; ok {
            ans = true
            break
        } else {
            m[n] = struct{}{}
        }
    }
    return ans
}
