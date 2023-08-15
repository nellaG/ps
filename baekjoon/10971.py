# https://www.acmicpc.net/problem/10971

from typing import List, Union

n = int(input())

# having cost of each city
matrix = [list(map(int, input().split())) for _ in range(n)]


def get_cost(perm):
    total_cost = 0
    for i in range(n - 1):
        from_idx = perm[i]
        to_idx = perm[i + 1]
        cost = matrix[from_idx][to_idx]
        if cost == 0:
            return False
        total_cost += cost

    final_cost = matrix[perm[n-1]][perm[0]]
    if final_cost == 0:
        return False
    total_cost += final_cost
    return total_cost


def next_permutation(perm) -> Union[List, None]:
    i = n - 1
    if perm[0] != 0:
        return None
    # 1
    while i > 0 and perm[i - 1] >= perm[i]:
        i -= 1
    if i <= 0:
        return None
    # 2
    j = n - 1
    while perm[j] <= perm[i - 1]:
        j -= 1
    # 3 (swap)
    perm[i - 1], perm[j] = perm[j], perm[i - 1]
    # 4
    idx = n - 1
    while i < idx:
        perm[i], perm[idx] = perm[idx], perm[i]
        i += 1
        idx -= 1

    return perm


def solution():
    perm = [x for x in range(n)]
    min_cost = 2100000000
    while True:
        if perm is None:
            break
        result = get_cost(perm)
        if result is not False:
            min_cost = min(min_cost, result)
        perm = next_permutation(perm)

    return min_cost


print(solution())
