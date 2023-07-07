# https://www.acmicpc.net/problem/1476

def solution():
    E, S, M = list(map(int, input().split(' ')))

    e = s = m = 1
    i = 1
    while True:
        if e == E and s == S and m == M:
            return i
        e += 1
        s += 1
        m += 1

        # reset
        if e == 16:
            e = 1
        if s == 29:
            s = 1
        if m == 20:
            m = 1
        i += 1


print(solution())
