m = []
for i in range(20):
    r = []
    for j in range(20):
        r.append(0)
    m.append(r)

N = int(input())
for i in range(1, N + 1):
    x, y = map(int, input().split())
    if i % 2 != 0:
        turn = 1
    else:
        turn = -1
    m[x][y] = turn
    res1, res2, res3, res4 = turn, turn, turn, turn
    # 1
    for j in range(1, 20):
        if y - j == 0:
            break
        if m[x][y - j] == turn:
            res1 += turn
        else:
            break
    for j in range(1, 20):
        if y + j == 20:
            break
        if m[x][y + j] == turn:
            res1 += turn
        else:
            break
    if abs(res1) == 5:
        print(i)
        break
    # 2
    for j in range(1, 20):
        if x - j == 0:
            break
        if m[x - j][y] == turn:
            res2 += turn
        else:
            break
    for j in range(1, 20):
        if x + j == 20:
            break
        if m[x + j][y] == turn:
            res2 += turn
        else:
            break
    if abs(res2) == 5:
        print(i)
        break
    # 3
    for j in range(1, 20):
        if x - j == 0 or y - j == 0:
            break
        if m[x - j][y - j] == turn:
            res3 += turn
        else:
            break
    for j in range(1, 20):
        if x + j == 20 or y + j == 20:
            break
        if m[x + j][y + j] == turn:
            res3 += turn
        else:
            break
    if abs(res3) == 5:
        print(i)
        break
    # 4
    for j in range(1, 20):
        if x + j == 20 or y - j == 0:
            break
        if m[x + j][y - j] == turn:
            res4 += turn
        else:
            break
    for j in range(1, 20):
        if x - j == 0 or y + j == 20:
            break
        if m[x - j][y + j] == turn:
            res4 += turn
        else:
            break
    if abs(res4) == 5:
        print(i)
        break

    if N == i:
        print(-1)
