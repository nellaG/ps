nums = [2, 7, 11, 15]
target = 9


def find_two_sum(nums, target):
    hashmap = {}  # 숫자가 키, 인덱스가 값
    for i, num in enumerate(nums):
        complement = target - num  # 차
        if complement in hashmap:
            return hashmap[complement], i
        hashmap[num] = i
    return None


print(find_two_sum(nums, target))
