# Runtime: 56 ms, faster than 99.21% of Python3 online submissions for Partition Array Into Three Parts With Equal Sum.
# Memory Usage: 17.6 MB, less than 93.13% of Python3 online submissions for Partition Array Into Three Parts With Equal Sum.
# 因为没有仔细审题，题目要求是“按照原数组的顺序”切分成三个等和片段，导致没有解出来
# 下面是评论区我认为最好的答案

class Solution:
    def canThreePartsEqualSum(self, A: List[int]) -> bool:
        s = sum(A)
        if s % 3 != 0:
            return False
        s /= 3
        # 因为没有对 acc 清零，所以第二次按顺序找到 三分之一的 s 之后 acc 应该是 2/3s
        targets = [2 * s, s]
        acc = 0
        for a in A:
            acc += a
            if acc == targets[-1]:
                targets.pop()
            if not targets:
                return True
        return False