# https://www.acmicpc.net/problem/6064


def solution():
    t = int(input())
    for _ in range(t):
        M, N, x, y = map(int, input().split())
        # 1씩 빼기
        # i : i % M, i % N

        x -= 1
        y -= 1
        ans = x
        while ans < N * M:
            if ans % N == y:
                print(ans + 1)
                break
            ans += M
        else:
            print(-1)

print(solution())
