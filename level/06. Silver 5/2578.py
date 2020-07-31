def numbering(d):
    for i in range(5):
        temp = list(map(int, input().split()))
        for j in range(5):
            d[temp[j]] = f'{i}{j}'


cs = {}
mc = {}
bingo = []

for _ in range(5):
    bingo.append([0, 0, 0, 0, 0])

numbering(cs)
numbering(mc)

count = 0
for i, v in mc.items():
    count += 1
    s = 0
    bingo[int(cs.get(i)[0])][int(cs.get(i)[1])] = 1
    if count >= 11:
        temp = 0
        for j in range(5):
            if sum(bingo[j]) == 5:
                s += 1
            temp = 0
            for k in range(5):
                if bingo[k][j] == 1:
                    temp += 1
            if temp == 5:
                s += 1

        temp = 0
        for j in range(5):
            if bingo[j][j] == 1:
                temp += 1
        if temp == 5:
            s += 1

        temp = 0
        k = 4
        for j in range(5):
            if bingo[j][k] == 1:
                temp += 1
            k -= 1
        if temp == 5:
            s += 1

    if s >= 3:
        print(count)
        break
