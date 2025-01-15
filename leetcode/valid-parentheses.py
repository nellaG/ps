class Solution:
    def isValid(self, s: str) -> bool:
        b_map = {")": "(", "}": "{", "]": "["}
        stack = []
        for x in s:
            if x in b_map.values():
                stack.append(x)
            else:  # close_b
                if len(stack) == 0 or b_map[x] != stack.pop():
                    return False
        return stack == []
