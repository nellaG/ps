from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        arr = ["", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]

        answer = []
        temp = []
        for d in digits:
            temp = []
            d = int(d)
            chars = list(arr[d])
            if answer == []:
                answer = chars
            else:
                for c in chars:
                    for a in answer:
                        temp.append(a + c)
                answer = temp
        return answer
