import sys

n, a, b = map(int, sys.stdin.readline().split())
a, b = a - 1, b - 1
x = []
f = 0
for i in range(n):
    x.append(list(map(int, sys.stdin.readline().split())))

z = x[a][b]

for i in range(n):
    for j in range(n):
        if i != a and j == b or i == a:
            if x[i][j] > z:
                print("ANGRY")
                f = 1
                break
    if f:
        break
    if i == n - 1:
        print("HAPPY")
