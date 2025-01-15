from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        answer = [nums[0]]
        for n in nums:
            if answer[-1] != n:
                answer.append(n)
            else:
                continue
        nums[:] = answer
        return len(answer)
