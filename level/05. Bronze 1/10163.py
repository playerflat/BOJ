a = []
for _ in range(102):
    a.append([0 for i in range(102)])

N = int(input())

for i in range(1, N + 1):
    x1, y1, x2, y2 = map(int, input().split())
    for x in range(x1, x1 + x2):
        for y in range(y1, y1 + y2):
            a[x][y] = i

for i in range(1, N + 1):
    count = 0
    for x in range(0, 102):
        for y in range(0, 102):
            if a[x][y] == i:
                count += 1
    print(count)