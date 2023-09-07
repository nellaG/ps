# https://www.acmicpc.net/problem/1309

# d[i][j] = i번째 줄 j번째 칸(1, 2) 에 사자를 놓는 경우의 수 j=0 사자 x
# d[i][0] =  d[i-1][1] + d[i-1][2]
# d[i][1] =  d[i-1][0] + d[i-1][2]
# d[i][2] =  d[i-1][0] + d[i-1][1]


n = int(input())
d = [[0] * 3 for _ in range(n)]  # not [[0] * 3] * n (reference)
mod = 9901

def solution():
    d[0][0] = 1
    d[0][1] = 1
    d[0][2] = 1
    for i in range(1, n):
        d[i][0] = (d[i-1][0] + d[i-1][1] + d[i-1][2]) % mod
        d[i][1] = (d[i-1][0] + d[i-1][2]) % mod
        d[i][2] = (d[i-1][0] + d[i-1][1]) % mod

    return sum(d[n-1]) % mod


print(solution())
