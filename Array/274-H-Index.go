// 暂时没有找到 golang 原生直接能倒序排序的简单方式
// 所以就用了正序排序然后倒序遍历的方式
// citations[i] >= len(citations)-i 这句是关键
// 倒序的来看，元素的值大于倒数的序号就带变有该"序号"个元素满足条件
func hIndex(citations []int) int {
    var ans int
    sort.Ints(citations)
    for i := len(citations)-1; i >=0; i-- {
        if citations[i] >= len(citations)-i {
            ans = len(citations)-i
        }
    }
    return ans
}
