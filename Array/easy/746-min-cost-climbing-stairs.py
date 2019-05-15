class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        f1, f2 = 0, 0
        for i in range(1, len(cost)):
            # cost 中的元素代表从第 i 个台阶爬上去(爬一个或者两个台阶的 cost 相同)的 cost
            # f2 + cost[i] 从前一个台阶再爬一个台阶上来
            # f1 + cost[i-1] 从前两个台阶爬两个台阶上来
            f1, f2 = f2, min(f2 + cost[i], f1 + cost[i-1])
        return f2