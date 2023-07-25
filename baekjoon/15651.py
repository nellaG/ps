# https://www.acmicpc.net/problem/15651

mlist = [0] * 9
nlist = [0] * 9

def solution():

    n, m = list(map(int, input().split()))

    return func(0, n, m)


def func(index: int, n, m):
    if index == m:
        ans = ""
        for x in nlist:
            if x == 0:
                break
            ans += f'{x} '
        print(ans)
        return
    for i in range(1, n + 1):
        #if mlist[i]:  중복선택 방지 코드
        #    continue
        mlist[i] = True
        nlist[index] = i
        func(index + 1, n, m)
        mlist[i] = False


solution()
