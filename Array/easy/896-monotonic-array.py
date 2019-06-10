# Runtime: 80 ms, faster than 94.34% of Python3 online submissions for Monotonic Array.
# Memory Usage: 17.5 MB, less than 38.61% of Python3 online submissions for Monotonic Array.

class Solution:
    def isMonotonic(self, A: List[int]) -> bool:
        if len(A) == 1:
            return True
        ld = 0
        for i in range(1, len(A)):
            if A[i - 1] == A[i]:
                continue
            diff = A[i] - A[i - 1]
            if ld * diff < 0:
                return False
            ld = diff
        return True