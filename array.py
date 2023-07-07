from collections import deque

def solution(rc, operations):

    for op in operations:
        if op == 'ShiftRow':
            rc = shift_row(rc)
        elif op == 'Rotate':
            rc = rotate(rc)
        else:
            rc = rc

    return rc


def rotate(mat):
    row = len(mat)
    col = len(mat[0])
    # topleft, topright, bottomleft, bottomright
    edge00 = edge01 = edge10 = edge11 = None

    edge00 = mat[1][0]
    edge01 = mat[0][-2]
    edge10 = mat[-1][1]
    edge11 = mat[-2][-1]

    # ->
    mat[0][1:-1] = mat[0][0:-2]

    # |
    # v
    for r in range(1, row - 1):

        mat[r][-1] = mat[r-1][-1]
        rev_r = row - r - 1
        mat[rev_r][0] = mat[rev_r+1][0]
    # <-
    mat[-1][1:-1] = mat[-1][2:]


    # calibration
    mat[0][0] = edge00
    mat[0][-1] = edge01
    mat[-1][0] = edge10
    mat[-1][-1] = edge11

    return mat


def shift_row(mat):
    q = deque(mat)
    last = q.pop()
    q.appendleft(last)
    return list(q)


arr = [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12]
]
op = ["ShiftRow", "Rotate", "ShiftRow", "Rotate"]
op = ["Rotate"]
res = solution(arr, op)
print(res)
