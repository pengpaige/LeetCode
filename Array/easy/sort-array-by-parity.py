class Solution:
    def sortArrayByParity(self, A: List[int]) -> List[int]:
        l, r = 0, len(A) - 1
        tmp = 0
        while l < r:
            if A[l] % 2 == 0:
                l += 1
                continue
            if A[r] % 2 != 0:
                r -= 1
                continue
            tmp = A[l]
            A[l] = A[r]
            A[r] = tmp
        return A