# https://www.acmicpc.net/problem/18290


n, m, k = list(map(int, input().split()))
a = [list(map(int, input().split())) for _ in range(n)]
c = [[False] * m for _ in range(n)]
ans = -2100000000
dx = [0, 0, 1, -1]
dy = [1, -1, 0, 0]


def go(px, py, cnt, s):
    if cnt == k:
        global ans
        if ans < s:
            ans = s
        return

    for x in range(px, n):
        for y in range(py if x == px else 0, m):
            if c[x][y]:
                continue  # 선택한 칸
            ok = True

            # 인접한 부분 검사
            for i in range(4):
                if not ok:
                    break
                nx, ny = x + dx[i], y + dy[i]  # 인접
                if 0 <= nx < n and 0 <= ny < m:
                    if c[nx][ny]:
                        ok = False

            if ok:
                c[x][y] = True
                go(x, y, cnt + 1, s + a[x][y])
                c[x][y] = False  # 현재 선택한 게 최대가 아닐 수도 있어서 원래 상태로 돌려놓아야 함


go(0, 0, 0, 0)
print(ans)
