# Runtime: 80 ms, faster than 50.89% of Python3 online submissions for Fair Candy Swap.
# Memory Usage: 15.4 MB, less than 24.37% of Python3 online submissions for Fair Candy Swap.

# 比较直接的思维方式，如果考虑节省空间可以使用 set 去重，然后用 set 里的元素值加上差值去做
class Solution:
    def fairCandySwap(self, A: List[int], B: List[int]) -> List[int]:
        dictA = {a: True for a in A}
        dictB = {b: True for b in B}
        semiDiff = (sum(A) - sum(B)) // 2
        for k in dictB:
            if k + semiDiff in dictA:
                return [k + semiDiff, k]