# https://www.acmicpc.net/problem/6603

from typing import Union, List

def next_permutation(perm) -> Union[List, None]:
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
    # 3 (swap)
    perm[i - 1], perm[j] = perm[j], perm[i - 1]
    # 4
    idx = n - 1
    while i < idx:
        perm[i], perm[idx] = perm[idx], perm[i]
        i += 1
        idx -= 1

    return perm


while True:
    n, *k = list(map(int, input().split()))
    if n == 0:
        break

    d = [0] * (n - 6) + [1] * 6  # pick 6

    answers = []
    while True:
        current = [k[i] for i in range(n) if d[i] == 1] # picked numbers
        answers.append(current)
        if next_permutation(d) is None:
            break

    answers.sort()
    for a in answers:
        print(' '.join(map(str, a)))

    print()
