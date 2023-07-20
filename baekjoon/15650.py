# https://www.acmicpc.net/problem/15650


nlist = [0] * 9


def solution():

    n, m = list(map(int, input().split()))

    return func(0, 0, n, m)


def func(index: int, start, n, m):
    if index == m:
        ans = []
        for x in nlist:
            if x == 0:
                break
            ans.append(str(x))
        if ans != []:
            print(' '.join(ans))
        return
    for i in range(start, n + 1):
        nlist[index] = i
        func(index + 1, i + 1, n, m)


solution()
