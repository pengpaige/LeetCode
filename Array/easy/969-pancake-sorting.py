# Input
# [3,2,4,1]
# Output
# [3,4,2,3,3,2,4,1]
# Expected
# [4,2,4,3]
# Time Submitted Status Runtime Memory Language
# a few seconds ago	Wrong Answer	N/A	N/A	python3
# 2 minutes ago	Wrong Answer	N/A	N/A	python3

# 题干里说的最大翻转次数不超过 10 倍的数组长度就认为答案正确，但实际上只接受一种答案，下面并不是官方接受的答案

class Solution:
    def pancakeSort(self, A: List[int]) -> List[int]:
        ret = list()
        length = len(A)
        sortedIndex = sorted(range(1, length+1), key=lambda i: -A[i-1])
        for index in sortedIndex:
            for flip in ret:
                if index < flip:
                    index = flip - index + 1
            ret.extend([index, length])
            length -= 1
        return ret