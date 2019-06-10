# 直接想到的方法
# Runtime: 720 ms, faster than 5.05% of Python3 online submissions for Move Zeroes.
# Memory Usage: 14.5 MB, less than 31.02% of Python3 online submissions for Move Zeroes.

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        zeroLoc = list()
        for i, v in enumerate(nums):
            if v == 0:
                zeroLoc.append(i)
        for i in range(len(zeroLoc)):
            # 每次移动之后原来数组中 0 的位置会向后移动 i 个位置
            for j in range(zeroLoc[i] - i, len(nums) - 1):
                nums[j], nums[j + 1] = nums[j + 1], nums[j]

#  上面的方法会对已经放到最后的 0 重复交换，下面是优化的方法
# Runtime: 652 ms, faster than 5.05% of Python3 online submissions for Move Zeroes.
# Memory Usage: 14.4 MB, less than 72.53% of Python3 online submissions for Move Zeroes.

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        zeroLoc = list()
        for i, v in enumerate(nums):
            if v == 0:
                zeroLoc.append(i)
        for i in range(len(zeroLoc)):
            # 最后 i 个已经是 0 的元素不再交换
            for j in range(zeroLoc[i] - i, len(nums) - i - 1):
                nums[j], nums[j + 1] = nums[j + 1], nums[j]

# 上面的优化效果不够明显
# Runtime: 48 ms, faster than 89.81% of Python3 online submissions for Move Zeroes.
# Memory Usage: 14.6 MB, less than 15.40% of Python3 online submissions for Move Zeroes.

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        zero = 0  # records the position of "0"
        for i in range(len(nums)):
            # 如果前几个数字都不是 0 的话，会产生相应数量的无效交换（ i == zero ）
            if nums[i] != 0:
                nums[i], nums[zero] = nums[zero], nums[i]
                zero += 1

# 针对上面方法无效交换问题的优化，使用空间少了，但是耗时反而长
# Runtime: 52 ms, faster than 52.07% of Python3 online submissions for Move Zeroes.
# Memory Usage: 14.2 MB, less than 99.02% of Python3 online submissions for Move Zeroes.

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        zero = 0  # records the position of "0"
        for i in range(len(nums)):
            if nums[i] != 0:
                if zero != i:
                    nums[i], nums[zero] = nums[zero], nums[i]
                zero += 1