# Runtime: 60 ms, faster than 99.54% of Python3 online submissions for Max Consecutive Ones.
# Memory Usage: 13.3 MB, less than 55.46% of Python3 online submissions for Max Consecutive Ones.

class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        curLen, maxLen = 0, 0
        for i in nums:
            if i == 0:
                curLen = 0
            else:
                curLen += 1
            if curLen > maxLen:
                maxLen = curLen
        return maxLen