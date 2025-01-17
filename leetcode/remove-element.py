from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        answer = []
        for n in nums:
            if val != n:
                answer.append(n)
            else:
                continue
        nums[:] = answer
        return len(answer)
