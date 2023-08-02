# https://www.acmicpc.net/problem/10972

n = int(input())

perm = list(map(int, input().split()))

def solution(perm):
    i = n - 1

    # 1
    while i > 0 and perm[i - 1] >= perm[i]:
        i -= 1
    if i <= 0:
        return False

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

    return True


result = -1 if not solution(perm) else ' '.join(list(map(str, perm)))
print(result)
