# https://www.acmicpc.net/problem/10819

from typing import List, Union

n = int(input())

perm = list(map(int, input().split()))
perm.sort()


def calc(perm):
    val = 0
    for i in range(n - 1):
        val += abs(perm[i] - perm[i + 1])
    return val


def next_permutation(perm) -> Union[List[int], None]:
    i = n - 1
    # 1
    while i > 0 and perm[i - 1] >= perm[i]:
        i -= 1
    if i <= 0:
        return None
    # 2
    j = n - 1
    while perm[j] <= perm[i - 1]:
        j -= 1
    # 3
    perm[i - 1], perm[j] = perm[j], perm[i - 1]
    # 4
    idx = n - 1
    while i < idx:
        perm[i], perm[idx] = perm[idx], perm[i]
        i += 1
        idx -= 1

    return perm


def solution():
    max_val = calc(perm)  # initial max value
    while next_permutation(perm) is not None:
        if (next_max_val := calc(perm)) > max_val:
            max_val = next_max_val
    return max_val


print(solution())
