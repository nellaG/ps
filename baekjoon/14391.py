# https://www.acmicpc.net/problem/14391

n, m = map(int, input().split())
a = [list(map(int, list(input()))) for _ in range(n)]

def solution():
    ans = 0
    for s in range(1 << (n * m)): # 0 -, 1 |  # 1-d bit
        # s: 모든 비트마스크의 경우
        sum_ = 0
        for i in range(n):  # row
            cur = 0
            for j in range(m):
                k = i * m + j
                if ((1 << k) & s) == 0:  # if k is horizontal
                    cur = cur * 10 + a[i][j]
                else:  # vertical
                    sum_ += cur
                    cur = 0
            sum_ += cur  # add last if last is horizontal (if vertical cur is 0)
        for j in range(m): # column
            cur = 0
            for i in range(n):
                k = i * m + j
                if ((1 << k) & s) != 0:  # if k is vertical
                    cur = cur * 10 + a[i][j]
                else:  # vertical
                    sum_ += cur
                    cur = 0
            sum_ += cur  # add last if last is vertical (if horizontal cur is 0)
        ans = max(sum_, ans)
    return ans


print(solution())



