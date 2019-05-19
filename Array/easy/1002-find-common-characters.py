from collections import Counter

class Solution:
    def commonChars(self, A: List[str]) -> List[str]:
        mainCounter = Counter(A[0])
        for str in A[1:]:
            mainCounter &= Counter(str)
        return list(mainCounter.elements())