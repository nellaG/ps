

# ["AN", "CF", "MJ", "RT", "NA"]	[5, 3, 2, 7, 5]	"TCMA"
# ["TR", "RT", "TR"]	[7, 1, 3]	"RCJA"
# 1 < 앞 캐릭터에 3점 7 -> 뒷캐릭터 3점

# R / T,  C / F,  J / M, A / N

def solution(survey, choices):
    answer = ''

    board = {
        'R': 0,
        'T': 0,
        'C': 0,
        'F': 0,
        'J': 0,
        'M': 0,
        'A': 0,
        'N': 0
    }

    for sur, choice in zip(survey, choices):
        front, back = sur
        score = 4 - choice
        if score < 0:
            board[back] += abs(score)
        elif score > 0:
            board[front] += abs(score)
    xx = iter(board.items())
    board = [(x, next(xx)) for x in xx]
    for front, back in board:
        char = front[0] if front[1] >= back[1] else back[0]
        answer += char
    return answer


solution(["AN", "CF", "MJ", "RT", "NA"], [5, 3, 2, 7, 5])  # "TCMA"
