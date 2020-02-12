// 类似大于一半的众数思想
// 出现一个或者两个候选人减为零的话，需要在下次循环里重置 cand 和 count
// 另外，因为有两个候选人，每次遍历到的数字最多跟一个候选人相等，两个候选人等于 n 互斥
// 最后，两个候选人都不等于 n 和上面一条中 有一个候选人等于 n 互斥
// 最最后，根据摩尔选举算法，最后需要验证一下最后胜出的候选人是否真的超过相应的票数
// 而且两个候选人不能同时将计数归零，否则可能出现候选人相等的情况
// 而且而且，必须先增加候选人票数再判断票数是否为零，否则也可能出现候选人相等的情况
func majorityElement(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    ans := make([]int, 0)
    cand1, cand2, count1, count2 := 0, 0, 0, 0
    for _, n := range nums {
        if cand1 == n {
            count1++
        } else if cand2 == n {
            count2++
        } else if count1 == 0 {
            cand1 = n
            count1 = 1
        } else if count2 == 0 {
            cand2 = n
            count2 = 1
        } else if cand1 != n && cand2 != n {
            count1--
            count2--
        }
    }
    
    if count1 > 0 {
        count1 = 0
        for _, n := range nums {
            if cand1 == n {
                count1++
            }
        }
        if count1 > len(nums)/3 {
            ans = append(ans, cand1)
        }
    }
    
    if count2 > 0 {
        count2 = 0
        for _, n := range nums {
            if cand2 == n {
                count2++
            }
        }
        if count2 > len(nums)/3 {
            ans = append(ans, cand2)
        }
    }
    
    return ans
}
