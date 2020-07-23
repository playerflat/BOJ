room = []
for i in range(int(input())):
    room.append(input())
x, y = 0, 0
for i in range(len(room)):

    temp = 0
    for j in range(len(room)):
        if room[i][j] == '.':
            temp += 1
        else:
            temp = 0
        if temp == 2:
            x += 1

    temp = 0
    for j in range(len(room)):
        if room[j][i] == '.':
            temp += 1
        else:
            temp = 0
        if temp == 2:
            y += 1

print(x, y)
