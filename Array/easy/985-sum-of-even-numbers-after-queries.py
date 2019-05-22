# O(n^2) 解法 会超时
class Solution:
    def sumEvenAfterQueries(self, A: List[int], queries: List[List[int]]) -> List[int]:
        ret = list()
        for i in range(len(queries)):
            A[queries[i][1]] = queries[i][0] + A[queries[i][1]]
            s = 0
            for a in A:
                if a % 2 == 0:
                    s = s + a
            ret.append(s)
        return ret

# O(N) 解法
# 每次求和可以看做一个连续的过程，本次对 A 中元素的改动对偶数和造成的影响是由于本次改动的元素造成的
# 奇数变成偶数需要加上新的值， 偶数变偶数需要减去旧值在加上新值，偶数变奇数需要减去旧值
# Runtime: 180 ms, faster than 62.45% of Python3 online submissions for Sum of Even Numbers After Queries.
# Memory Usage: 17.7 MB, less than 12.83% of Python3 online submissions for Sum of Even Numbers After Queries.
class Solution:
    def sumEvenAfterQueries(self, A: List[int], queries: List[List[int]]) -> List[int]:
        ret = list()
        s = 0
        for a in A:
            if a % 2 == 0:
                s = s + a
        for i in range(len(queries)):
            old = A[queries[i][1]]
            new = queries[i][0] + A[queries[i][1]]
            if old % 2 == 0:
                if new % 2 != 0:
                    s = s - old
                else:
                    s = s - old + new
            else:
                if new % 2 == 0:
                    s = s + new
            A[queries[i][1]] = new
            ret.append(s)
        
        return ret