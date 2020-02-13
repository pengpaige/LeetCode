/**
 * @param words: a list of words
 * @param word1: a string
 * @param word2: a string
 * @return: the shortest distance between word1 and word2 in the list
 */
 
// 找到两个单词所有的位置，分别存储到两个 slice
// 嵌套遍历两个 slice 找差最小的情况
func shortestDistance (words []string, word1 string, word2 string) int {
    // Write your code here
    ans := len(words)-1
    w1idx := make([]int, 0) 
    w2idx := make([]int, 0)
    for i, n := range words {
        if n == word1 {
            w1idx = append(w1idx, i)
        } else if n == word2 {
            w2idx = append(w2idx, i)
        }
    }
    for _, m := range w1idx {
        for _, n := range w2idx {
            diff := m-n
            if diff < 0 {
                diff *=-1
            }
            if diff < ans {
                ans = diff 
            }
        }
    }
    return ans
}

