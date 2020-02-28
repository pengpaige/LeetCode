// 利用归并排序的回溯流程对未排序的部分中的逆序对进行查找
// 另外还有一个技巧
// 已排序的 l 和 r 两个 slice 中如果l中的 i 满足逆序对的条件，那么 i 后面的元素也都满足
func reversePairs(nums []int) int {
    var ans int
    mergeCount(nums, 0, len(nums)-1, &ans)
    return ans
}

func mergeCount(nums []int, start, end int, count *int) []int {
    if start > end {
        return nil
    }
    if start == end {
        return []int{nums[start]}
    }
    mid := start + (end - start)/2
    l := mergeCount(nums, start, mid, count)
    r := mergeCount(nums, mid+1, end, count)
    return merge(l, r, count)
}

func merge(l, r []int, count *int) []int {
    var i, j int
    ret := make([]int, 0)
    for i < len(l) && j < len(r) {
        if l[i] > 2*r[j] { // 这只是逆序对的判断依据不是归并的交换条件
            *count += len(l)-i // len(l)-1-i+1
            j++
        } else {
            i++
        }
    }
    i, j = 0, 0
    for i < len(l) && j < len(r) {
        if l[i] >= r[j] {
            ret = append(ret, r[j])
            j++
        } else {
            ret = append(ret, l[i])
            i++
        }
    }
    for i < len(l) {
        ret = append(ret, l[i])
        i++
    }
    for j < len(r) {
        ret = append(ret, r[j])
        j++
    }
    return ret
}
