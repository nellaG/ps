# https://www.acmicpc.net/problem/15661

N = int(input())

s = [list(map(int, (input().split()))) for _ in range(N)]


def go(index, team_1, team_2):

    # 정답일때
    if index == N:
        if (len(team_1) < 1) or (len(team_2) < 1):
            return -1

        t1 = 0
        t2 = 0

        for p1 in team_1:
            for p2 in team_1:
                if p1 == p2:
                    continue
                t1 += s[p1][p2]
        for p1 in team_2:
            for p2 in team_2:
                if p1 == p2:
                    continue
                t2 += s[p1][p2]
        diff = abs(t1 - t2)
        return diff

    ans = -1

    # 1번팀에 넣어보고 다시 빼기
    t1 = go(index + 1, team_1 + [index], team_2)
    if ans == -1 or (t1 != -1 and ans > t1):
        ans = t1

    # 2번팀에 넣어보고 다시 빼기
    t2 = go(index + 1, team_1, team_2 + [index])
    if ans == -1 or (t2 != -1 and ans > t2):
        ans = t2

    return ans


print(go(0, [], []))
