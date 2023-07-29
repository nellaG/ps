# https://www.acmicpc.net/problem/14889

N = int(input())

s = [list(map(int, (input().split()))) for _ in range(N)]


def go(index, team_1, team_2):

    # 정답일때
    if index == N:
        if (len(team_1) != N // 2) or (len(team_2) != N // 2):
            return -1

        t1 = 0
        t2 = 0

        for i in range(N // 2):
            for j in range(N // 2):
                if i == j:
                    continue
                t1 += s[team_1[i]][team_1[j]]
                t2 += s[team_2[i]][team_2[j]]

        diff = abs(t1 - t2)
        return diff

    # 정답 아닐때
    if (len(team_1) > N // 2) or (len(team_2) > N // 2):
        return -1
    # 아직일 때

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
