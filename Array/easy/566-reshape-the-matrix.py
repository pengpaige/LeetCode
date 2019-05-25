# Runtime: 88 ms, faster than 92.03% of Python3 online submissions for Reshape the Matrix.
# Memory Usage: 14.3 MB, less than 37.26% of Python3 online submissions for Reshape the Matrix.
class Solution:
    def matrixReshape(self, nums: List[List[int]], r: int, c: int) -> List[List[int]]:
        if r * c != len(nums) * len(nums[0]):
            return nums
        tmp = list()
        for row in nums:
            tmp.extend(row)
        ret = list()
        for i in range(r):
            l = tmp[i*c:(i+1)*c]
            ret.append(l)
            
        return ret


# 参考讨论区使用迭代器的方式
# Runtime: 88 ms, faster than 92.03% of Python3 online submissions for Reshape the Matrix.
# Memory Usage: 14.4 MB, less than 17.14% of Python3 online submissions for Reshape the Matrix.
class Solution:
    def matrixReshape(self, nums: List[List[int]], r: int, c: int) -> List[List[int]]:
        if r * c != len(nums) * len(nums[0]):
            return nums
        tmp = iter(val for row in nums for val in row)
        return [[next(tmp) for c in range(c)] for r in range(r)]