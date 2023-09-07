# https://www.acmicpc.net/problem/14501

n = int(input())
pair = [list(map(int, (input().split()))) for _ in range(n)]
answer = 0
t = [0] * (n + 1)
p = [0] * (n + 1)
inf = -10**9
d = [-1] * (n + 1)


def solution(day):
    if day == n + 1:
        return 0
    if day > n + 1:
        return inf
    if d[day] != -1:
        return d[day]
    t1 = solution(day + 1)
    t2 = pair[day-1][1] + solution(day + pair[day-1][0])
    d[day] = max(t1, t2)
    return d[day]


print(solution(1))


