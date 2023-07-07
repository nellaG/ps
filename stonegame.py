# max score

from functools import lru_cache

class Solution:
    def stoneGameV(self, stone) -> int:
        prefix = [0]

        for i in range(len(stone)):
            prefix += [prefix[-1] + stone[i]]
        print(prefix)

        @lru_cache()
        def dfs(left, right):
            if right - left == 1:
                return 0
            maxval = -1
            for i in range(left + 1, right):
                leftsum = prefix[i] - prefix[left]
                rightsum = prefix[right] - prefix[i]

                if leftsum > rightsum:
                    maxval = max(maxval, rightsum

        return stone


sol = Solution()

input_ = [6, 2, 3, 4, 5, 5]  # 18

input_ = [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]  # 37

print(sol.stoneGameV(input_))
