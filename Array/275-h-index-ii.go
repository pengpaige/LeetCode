/*
从测试例来看数组中是可能有重复数字的
二分查找 需要注意的三个点
1. 终止条件 l<=r
2. 左右移动 l r 要使用 mid+1 和 mid-1 不要直接使用 mid
3. 遍历元素非常多时需要考虑 l+r 的溢出问题，最好使用 l+(r-l)/2
*/
func hIndex(citations []int) int {
    var ans, mid int
    l, r, n := 0, len(citations)-1, len(citations)
    for l <= r {
        mid = (l+r)/2
        if citations[mid] < n-mid {
            l = mid+1
        } else if mid == 0 || citations[mid-1] < n-(mid-1) {
            ans = n-mid
            break
        } else { // mid != 0 && citations[mid-1] >= n-(mid-1)
            r = mid-1
        } 
    }
    return ans
}
