# https://www.acmicpc.net/problem/1759

L, C = list(map(int, input().split()))
alphabet = input().split()
alphabet.sort()


def go(n, alpha, pw, i):

    if len(pw) == n:
        if check(pw):
            print(pw)
            return
    if i >= len(alpha):
        return

    # 알파벳 쓴다면
    go(n, alpha, pw + alpha[i], i + 1)  # 다음 알파벳

    # 안쓴다면
    go(n, alpha, pw, i + 1)  # 다음 알파벳


def check(pw):
    ja = 0
    mo = 0
    for letter in pw:
        if letter in ['a', 'e', 'i', 'o', 'u']:
            mo += 1
        else:
            ja += 1
    return (mo >= 1 and ja >= 2)


go(L, alphabet, '', 0)
