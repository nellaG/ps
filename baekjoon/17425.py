# https://www.acmicpc.net/problem/17425
# run pypy3

def solution():
    answers = []
    MAX = 1000000
    g_n = [1] * (MAX + 1)
    s_i = [0] * (MAX + 1)  # s[i]  = g_n[1] _ ... g_n[i]
    for i in range(2, MAX + 1):
        for j in range(1, MAX + 1):
            if i * j > MAX:
                break
            g_n[i * j] += i
    for i in range(1, MAX + 1):
        s_i[i] = s_i[i-1] + g_n[i]

    count = int(input())
    for _ in range(count):
        num = int(input())
        answers.append(s_i[num])
    return '\n'.join(map(str, answers))


print(solution())
