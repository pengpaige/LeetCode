func canCompleteCircuit(gas []int, cost []int) int {
    if gas == nil || cost == nil || len(gas) == 0 || len(cost) == 0 {
        return -1
    }
    
    total, curr, start := 0, 0, 0
    for i := 0; i < len(gas); i++ {
        curr += gas[i] - cost[i]
        if curr < 0 {
            start = i+1
            total += curr
            curr = 0
        }
    }
    if total + curr >= 0 {
        return start
    } else {
        return -1
    }
}
