// 给定一个形如 “HH:MM” 表示的时刻，利用当前出现过的数字构造下一个距离当前时间最近的时刻。每个出现数字都可以被无限次使用。

// 你可以认为给定的字符串一定是合法的。例如，“01:34” 和 “12:09” 是合法的，“1:34” 和 “12:9” 是不合法的。

//  

// 样例 1:

// 输入: "19:34"
// 输出: "19:39"
// 解释: 利用数字 1, 9, 3, 4 构造出来的最近时刻是 19:39，是 5 分钟之后。结果不是 19:33 因为这个时刻是 23 小时 59 分钟之后。
//  

// 样例 2:

// 输入: "23:59"
// 输出: "22:22"
// 解释: 利用数字 2, 3, 5, 9 构造出来的最近时刻是 22:22。 答案一定是第二天的某一时刻，所以选择可构造的最小时刻。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/next-closest-time
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 穷举
func nextClosestTime(time string) string {
    bt := append([]byte(time[:2]), []byte(time[3:])...)
    st:= string(time[:2])+string(time[3:])
    sort.Slice(bt, func(i, j int) bool {
        return bt[i] < bt[j]
    })
    for i := range bt {
        for j := range bt {
            for k := range bt {
                if bt[k] > '5' {
                    continue
                }
                for l := range bt {
                    tmp := string(append([]byte{}, bt[i], bt[j], bt[k], bt[l]))
                    if tmp > st && tmp <= "2359" {
                        return string(tmp[:2])+":"+string(tmp[2:])
                    }
                }
            }
        }
	}
	// 如果穷举没有找到下一个最近的时刻
	// 说明下一个最近的时刻在第二天
	// 直接用 bt 中的最小非零字符组合成时间格式即可
    b := bt[0]
    if bt[0] == '0' {
        b = bt[1]
    }
    return string(append([]byte{}, b, b, ':', b, b))
}