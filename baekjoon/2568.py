# https://www.acmicpc.net/problem/2568
# n^2로 안되는 문제

n = int(input())

a = [list(map(int, input().split())) for _ in range(n)]
# d[i] = i번째까지 왔을떄 교차 안하는 (증가하는) 길이
#d = [1] * n
ans = set()
a = sorted(a, key=lambda x: x[0])


def solution():
    for i in range(n):
        for j in range(i):  # j < i
            if a[j][1] >= a[i][1]:
                ans.add(a[j][0])
    return len(ans)


print(solution())
print('\n'.join(map(str, ans)))
