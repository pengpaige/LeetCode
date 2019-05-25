# Runtime: 36 ms, faster than 99.33% of Python3 online submissions for Toeplitz Matrix.
# Memory Usage: 13.3 MB, less than 30.50% of Python3 online submissions for Toeplitz Matrix.

# 只比较左上角的元素是否和当前遍历到的元素相等就可以了，或者你也可以与右下角元素相比较
class Solution:
    def isToeplitzMatrix(self, matrix: List[List[int]]) -> bool:
        for r, row in enumerate(matrix):
            for c, val in enumerate(row):
                if r == 0 or c == 0:
                    continue
                if matrix[r-1][c-1] != matrix[r][c]:
                    return False
        return True