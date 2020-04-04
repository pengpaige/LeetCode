type Bucket struct {
    Max int
    Min int
}

func NewBucket() *Bucket{
    return &Bucket{
        Max: -1,
        Min: 1<<32-1, 
    }
}

func maximumGap(nums []int) int {
    l := len(nums)
    if l < 2 {
        return 0
    }
    
    maxNum, minNum := -1, 1<<32-1
    for _, n := range nums {
        maxNum = max(maxNum, n)
        minNum = min(minNum, n)
    }
    
    // 这里要注意，最小的桶间隔是 1 
    // 这里不设置最小桶间隔 1 的话当输入类似 [1,1,1,1] 这种的话会出现除以 0 的问题
    bktSize := max(1, (maxNum - minNum) / (l - 1))
    bktCount := (maxNum - minNum) / bktSize + 1
    bktList := make([]*Bucket, bktCount)
    
    for _, n := range nums {
        bktIndex := (n - minNum) / bktSize
        if bktList[bktIndex] == nil {
            bktList[bktIndex] = NewBucket()
        }
        bktList[bktIndex].Max = max(bktList[bktIndex].Max, n)
        bktList[bktIndex].Min = min(bktList[bktIndex].Min, n)
    }
    
    var preBkt *Bucket
    maxGap := 0
    for i := 0; i < bktCount; i++ {
        if bktList[i] == nil {
            continue
        }
        if preBkt == nil {
            preBkt = bktList[i]
            continue
        }
        maxGap = max(maxGap, bktList[i].Min - preBkt.Max)
        preBkt = bktList[i]
    }
    return maxGap
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
