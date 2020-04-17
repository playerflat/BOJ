import sys

a = int(sys.stdin.readline())
b = []
c = 2
for i in range(a):
    b.append(list(map(int, sys.stdin.readline().split())))
for i in range(1, a):
    for j in range(c):
        if j == 0:
            b[i][j] = b[i][j] + b[i - 1][j]
        elif i == j:
            b[i][j] = b[i][j] + b[i - 1][j - 1]
        else:
            b[i][j] = max(b[i - 1][j - 1], b[i - 1][j]) + b[i][j]
    c += 1
print(max(b[a - 1]))
