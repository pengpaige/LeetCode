class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        d = {}
        for i in range(len(numbers)):
            if numbers[i] in d:
                return d[numbers[i]], i + 1
            else:
                d[target - numbers[i]] = i + 1