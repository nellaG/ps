from typing import List


def trap(height: List[int]) -> int:
    length = len(height)
    if len == 0:
        return 0
    result = 0
    left_max = [0] * length
    right_max = [0] * length
    left_max[0] = height[0]
    right_max[-1] = height[-1]
    for i in range(1, length):
        left_max[i] = max(height[i], left_max[i - 1])

    for i in range(length - 2, -1, -1):
        print(i)
        right_max[i] = max(height[i], right_max[i + 1])
    print(left_max)
    print(right_max)
    print(height)
    for i in range(1, length - 1):
        result += min(left_max[i], right_max[i]) - height[i]

    return result


input_ = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
input_ = [4, 2, 0, 3, 2, 5]
print(trap(input_))
