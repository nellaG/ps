# https://www.acmicpc.net/problem/10974

n = int(input())

perm = [x for x in range(1, n + 1)]

def go(perm):
    i = n - 1

    # 1
    while i > 0 and perm[i - 1] >= perm[i]:
        i -= 1
    if i <= 0:
        return False

    # 2
    j = n - 1
    while perm[j] <= perm[i - 1]:
        j -= 1

    # 3
    perm[i - 1], perm[j] = perm[j], perm[i - 1]

    # 4
    idx = n - 1
    while i < idx:
        perm[i], perm[idx] = perm[idx], perm[i]
        i += 1
        idx -= 1

    return True


if __name__ == '__main__':
    if n == 1:
        print(1)
    else:
        print(' '.join(map(str, perm)))
        while True:
            result = go(perm)
            if result:
                print(' '.join(map(str, perm)))
            else:
                break
