# Runtime: 60 ms, faster than 87.51% of Python3 online submissions for Transpose Matrix.
# Memory Usage: 13.7 MB, less than 39.70% of Python3 online submissions for Transpose Matrix.
class Solution:
    def transpose(self, A: List[List[int]]) -> List[List[int]]:
        # 这是一个 py3 初始化多维数组的小技巧
        ret = [[ 0 for i in range(len(A))] for j in range(len(A[0]))]
        for i in range(len(A)):
            for j in range(len(A[0])):
                    ret[j][i] = A[i][j]
        return ret