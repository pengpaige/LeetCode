#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#
# https://leetcode.com/problems/two-sum/description/
#
# algorithms
# Easy (43.28%)
# Total Accepted:    1.7M
# Total Submissions: 3.9M
# Testcase Example:  '[2,7,11,15]\n9'
#
# Given an array of integers, return indices of the two numbers such that they
# add up to a specific target.
# 
# You may assume that each input would have exactly one solution, and you may
# not use the same element twice.
# 
# Example:
# 
# 
# Given nums = [2, 7, 11, 15], target = 9,
# 
# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].
# 
# 
# 
# 
#
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = {}
        for i in range(len(nums)):
            # 重复并且是 target 一半的元素直接返回
            if nums[i] in d:
                if 2 * nums[i] == target:
                    return [d[nums[i]], i]
            d[nums[i]] = i
        for i in d:
            # 这里得排除恰好是 target 一半的不重复的元素，避免后面 target - i 的时候得到本身
            if target == 2 * i:
                continue
            if target - i in d:
                return [d[i], d[target - i]]  

