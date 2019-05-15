class Solution:
    def climbStairs(self, n: int) -> int:
        if n == 1:
            return 1
        tmp = [-1 for i in range(n+1)]
        tmp[1], tmp[2] = 1, 2
        return self.dp(n, tmp)
        
    def dp(self, n, tmp):
        if tmp[n] != -1:
            return tmp[n]
        tmp[n] = self.dp(n - 1, tmp) + self.dp(n - 2, tmp)
        return tmp[n]