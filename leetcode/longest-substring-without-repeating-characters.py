class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        charmap = {}
        left = 0
        answer = 0

        for i, x in enumerate(s):
            if x in charmap and charmap[x] >= left:  # 중복
                left = charmap[x] + 1

            charmap[x] = i
            answer = max(answer, i - left + 1)

        return answer
